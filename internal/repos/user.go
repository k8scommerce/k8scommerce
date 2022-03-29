package repos

import (
	"context"
	"fmt"

	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func newUser(repo *repo) User {
	return &userRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type User interface {
	Exists() bool
	Deleted() bool
	Create(user *models.User) error
	Update(user *models.User) error
	Save() error
	Upsert() error
	Delete(id int64) error
	Login(username, password string) (res *models.User, err error)
	GetAllPermissionGroups(currentPage, pageSize int64, sortOn string) (res *getAllPermissionGroupsResponse, err error)
	GetAllUsers(currentPage, pageSize int64, sortOn string) (res *getAllUsersResponse, err error)
}

type userRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.User
}

type getAllPermissionGroupsResponse struct {
	PagingStats      PagingStats
	PermissionGroups []models.PermissionGroup
}

type getAllUsersResponse struct {
	PagingStats PagingStats
	Users       []models.User
}

func (m *userRepo) GetAllPermissionGroups(currentPage, pageSize int64, sortOn string) (res *getAllPermissionGroupsResponse, err error) {
	orderBy, err := BuildOrderBy(sortOn, map[string]string{
		"id":         "pg",
		"group_name": "pg",
	})
	if err != nil {
		return nil, err
	}

	// set a default order by
	if orderBy == "" {
		orderBy = "ORDER BY pg.group_name ASC"
	}
	offset := fmt.Sprintf("OFFSET %d", currentPage*pageSize)
	limit := fmt.Sprintf("LIMIT %d", pageSize)

	nstmt, err := m.db.PrepareNamed(fmt.Sprintf(`
			select 
				-- permissiongroup
				pg.id AS "permissiongroup.id",
				pg.group_name AS "permissiongroup.group_name",
				
				-- stats
				COUNT(pg.*) OVER() AS "pagingstats.total_records"
			from permission_group pg
			%s
			%s
			%s
		`, orderBy, offset, limit))
	if err != nil {
		return nil, fmt.Errorf("error::GetAllPermissionGroups::%s", err.Error())
	}

	var result []*struct {
		PermissionGroup models.PermissionGroup
		PagingStats     PagingStats
	}

	err = nstmt.Select(&result,
		map[string]interface{}{
			"offset":   currentPage * pageSize,
			"limit":    pageSize,
			"order_by": orderBy,
		})

	var users []models.PermissionGroup
	if len(result) > 0 {
		var stats *PagingStats
		for i, r := range result {
			if i == 0 {
				stats = r.PagingStats.Calc(pageSize)
			}
			users = append(users, r.PermissionGroup)
		}

		out := &getAllPermissionGroupsResponse{
			PermissionGroups: users,
			PagingStats:      *stats,
		}
		return out, err
	}

	return nil, err
}

func (m *userRepo) GetAllUsers(currentPage, pageSize int64, sortOn string) (res *getAllUsersResponse, err error) {
	orderBy, err := BuildOrderBy(sortOn, map[string]string{
		"first_name": "u",
		"last_name":  "u",
		"email":      "u",
	})
	if err != nil {
		return nil, err
	}

	// set a default order by
	if orderBy == "" {
		orderBy = "ORDER BY u.last_name ASC, u.first_name ASC"
	}
	offset := fmt.Sprintf("OFFSET %d", currentPage*pageSize)
	limit := fmt.Sprintf("LIMIT %d", pageSize)

	nstmt, err := m.db.PrepareNamed(fmt.Sprintf(`
			select 
				-- user
				u.id AS "user.id",
				u.first_name AS "user.first_name",
				u.last_name AS "user.last_name",
				u.email AS "user.email",
				u.password AS "user.password",
				
				-- stats
				COUNT(u.*) OVER() AS "pagingstats.total_records"
			from users u
			%s
			%s
			%s
		`, orderBy, offset, limit))
	if err != nil {
		return nil, fmt.Errorf("error::GetAllUsers::%s", err.Error())
	}

	fmt.Println(fmt.Sprintf(`
	select 
		-- user
		u.id AS "user.id",
		u.first_name AS "user.first_name",
		u.last_name AS "user.last_name",
		u.email AS "user.email",
		u.password AS "user.password",
		
		-- stats
		COUNT(u.*) OVER() AS "pagingstats.total_records"
	from users u
	%s
	%s
	%s
`, orderBy, offset, limit))

	var result []*struct {
		User        models.User
		PagingStats PagingStats
	}

	err = nstmt.Select(&result,
		map[string]interface{}{
			"offset":   currentPage * pageSize,
			"limit":    pageSize,
			"order_by": orderBy,
		})

	var users []models.User
	if len(result) > 0 {
		var stats *PagingStats
		for i, r := range result {
			if i == 0 {
				stats = r.PagingStats.Calc(pageSize)
			}
			users = append(users, r.User)
		}

		out := &getAllUsersResponse{
			Users:       users,
			PagingStats: *stats,
		}
		return out, err
	}

	return nil, err
}

func (m *userRepo) Login(username, password string) (res *models.User, err error) {
	res, err = models.UserByEmail(m.ctx, m.db, username)
	if err != nil {
		return nil, err
	}

	if m.checkPasswordHash(password, res.Password) {
		return res, nil
	}

	return nil, fmt.Errorf("error: incorrect username and password combination")
}

func (m *userRepo) Create(user *models.User) error {
	// hash the password
	hash, _ := m.hashPassword(user.Password)
	user.Password = hash

	if err := user.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *userRepo) Update(user *models.User) error {
	if user.ID == 0 {
		return fmt.Errorf("error: can't update user, missing user ID")
	}
	if err := user.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *userRepo) Save() error {
	return m.User.Save(m.ctx, m.db)
}

func (m *userRepo) Upsert() error {
	return m.User.Upsert(m.ctx, m.db)
}

func (m *userRepo) Delete(id int64) error {
	user, err := models.ProductByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return user.Delete(m.ctx, m.db)
}

func (m *userRepo) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (m *userRepo) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
