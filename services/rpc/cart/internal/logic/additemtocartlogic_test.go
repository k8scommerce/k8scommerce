package logic_test

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/cart/internal/config"
	"k8scommerce/services/rpc/cart/internal/logic"
	"k8scommerce/services/rpc/cart/internal/svc"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/localrivet/gcache"

	"github.com/iancoleman/strcase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tal-tech/go-zero/core/conf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ = Describe("CartLogic", func() {
	defer GinkgoRecover()

	var testProduct *models.Product
	var testVariant *models.Variant
	var testPrice *models.Price
	var c config.Config
	var storeId int64 = 1
	var customerId int64 = 1

	conf.MustLoad("../../etc/cart.yaml", &c)
	svcCtx := svc.NewServiceContext(c)
	universe := gcache.NewUniverse(c.ListenOn)
	ctx := context.Background()

	expectNilErr := func(err error) {
		Expect(err).To(BeNil())
	}

	productName := "Test Product"
	productResponse, err := svcCtx.Repo.Product().GetProductBySlug(storeId, strcase.ToKebab(productName))
	expectNilErr(err)
	if productResponse != nil {
		svcCtx.Repo.CartItem().Delete(customerId, strcase.ToKebab(productName), true)
		err = svcCtx.Repo.Product().Delete(productResponse.Product.ID)
		expectNilErr(err)
	}

	testProduct = &models.Product{
		Slug:             strcase.ToKebab(productName),
		Name:             productName,
		ShortDescription: sql.NullString{String: "Short description", Valid: true},
		Description:      sql.NullString{String: "Longer description", Valid: true},
		MetaTitle:        sql.NullString{String: productName, Valid: true},
		MetaDescription:  sql.NullString{String: "Short description", Valid: true},
		MetaKeywords:     sql.NullString{String: strings.Join(strings.Split("Short description", " "), ","), Valid: true},
		Promotionable:    true,
	}
	if err := testProduct.Insert(context.Background(), svcCtx.Repo.GetRawDB()); err == nil {
		// add a variant
		testVariant = &models.Variant{
			ProductID:      testProduct.ID,
			IsDefault:      true,
			Sku:            testProduct.Slug,
			CostAmount:     sql.NullInt64{Int64: int64(utils.RandInt(999, 10999)), Valid: true},
			CostCurrency:   sql.NullString{String: "USD", Valid: true},
			TrackInventory: true,
			SortOrder:      1,
		}
		if err := testVariant.Insert(context.Background(), svcCtx.Repo.GetRawDB()); err == nil {
			money1 := float64(testVariant.CostAmount.Int64) * utils.RandFloat(1.2, 3)
			money2 := float64(testVariant.CostAmount.Int64) * utils.RandFloat(3, 3.4)

			testPrice = &models.Price{
				VariantID:       testVariant.ID,
				Amount:          int64(money1 * 100),
				Currency:        testVariant.CostCurrency,
				CompareAtAmount: sql.NullInt64{Int64: int64(money2 * 100), Valid: true},
			}
			if err := testPrice.Insert(context.Background(), svcCtx.Repo.GetRawDB()); err == nil {
			} else {
				expectNilErr(err)
			}
		} else {
			expectNilErr(err)
		}
	} else {
		expectNilErr(err)
	}
	productResponse, err = svcCtx.Repo.Product().GetProductBySlug(storeId, strcase.ToKebab(productName))
	expectNilErr(err)

	for _, v := range productResponse.Variants {
		if v.IsDefault {
			testVariant = &v
			break
		}
	}

	Expect(testProduct).To(Not(BeNil()))
	Expect(testVariant).To(Not(BeNil()))
	Expect(testPrice).To(Not(BeNil()))

	// create a cart
	cartResponse, err := svcCtx.Repo.Cart().Upsert(&models.Cart{
		CustomerID: customerId,
	})
	expectNilErr(err)
	Expect(cartResponse).ToNot(BeNil())

	Describe("AddItemToCartLogic", func() {
		var addItemToCartLogic *logic.AddItemToCartLogic
		// var removeItemToCartLogic *logic.RemoveItemInCartLogic
		BeforeEach(func() {
			addItemToCartLogic = logic.NewAddItemToCartLogic(ctx, svcCtx, universe)
			// removeItemToCartLogic = logic.NewRemoveItemInCartLogic(ctx, svcCtx, universe)

			// out, err := removeItemToCartLogic.RemoveItemInCart(&cart.RemoveItemInCartRequest{
			// 	CustomerId: int64(customerId),
			// 	CartId: int64(cartId),
			// 	Sku:    testVariant.Sku,
			// })

			// expectNilErr(err)
			// Expect(out).To(Not(BeNil()))
		})

		It("should return a valid addItemToCartLogic", func() {
			Expect(addItemToCartLogic).To(Not(BeNil()))
		})

		Describe("AddItemToCart", func() {
			expires := time.Now().Add(time.Hour * 1)

			It("should add an item to cart", func() {
				// AddItemToCart(in *cart.AddItemToCartRequest) (*cart.AddItemToCartResponse, error)
				request := &cart.AddItemToCartRequest{
					CustomerId: int64(customerId),
					Item: &cart.Item{
						CustomerId: int64(customerId),
						Sku:        testVariant.Sku,
						Quantity:   1,
						Price:      testPrice.Amount,
						ExpiresAt:  timestamppb.New(expires),
					},
				}

				response, err := addItemToCartLogic.AddItemToCart(request)
				// fmt.Printf("response: %#v", response)
				expectNilErr(err)
				Expect(response).To(Not(BeNil()))
			})
		})
	})
})
