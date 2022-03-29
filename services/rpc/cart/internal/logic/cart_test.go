package logic_test

import (
	"context"
	"k8scommerce/internal/session"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Createcartlogic", func() {
	var cartId string
	var sessionId string

	var createXSessions = func(num int) {
		for i := 0; i < num; i++ {
			in := &cart.CreateCartRequest{
				StoreId: 1,
			}
			ctx := context.Background()
			_, err := srv.CreateCart(ctx, in)
			Expect(err).To(BeNil())
		}
	}

	Describe("Build Cart With Customer & Items", func() {
		truncateCartItem()
		truncateCart()

		Describe("CreateCartLogic", func() {
			It("should create a cart", func() {
				createXSessions(10)

				in := &cart.CreateCartRequest{
					StoreId: 1,
				}
				ctx := context.Background()
				response, err := srv.CreateCart(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				cartId = response.Cart.Id
				Expect(cartId).ToNot(Equal(0))
			})
		})

		Describe("GetByCartId", func() {
			It("should get a cart from the database", func() {
				in := &cart.GetByCartIdRequest{
					CartId: cartId,
				}
				ctx := context.Background()
				response, err := srv.GetByCartId(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				sessionId = response.SessionId
			})
		})

		Describe("GetBySessionId", func() {
			It("should get a cart from cache", func() {
				in := &cart.GetBySessionIdRequest{
					SessionId: sessionId,
					CartId:    cartId,
				}
				ctx := context.Background()
				response, err := srv.GetBySessionId(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				sessionId = response.SessionId
			})

			It("should return a valid cart if cache fails", func() {
				in := &cart.GetBySessionIdRequest{
					SessionId: session.NewSessionId(),
					CartId:    cartId,
				}
				ctx := context.Background()
				response, err := srv.GetBySessionId(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
			})

			It("should fail getting an invalid cart from cache", func() {
				in := &cart.GetBySessionIdRequest{
					SessionId: session.NewSessionId(),
					CartId:    uuid.NewString(),
				}
				ctx := context.Background()
				response, err := srv.GetBySessionId(ctx, in)
				Expect(err).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
		})

		Describe("AddItemLogic", func() {
			It("should add an item to cart", func() {
				in := &cart.AddItemRequest{
					CartId: cartId,
					Item: &cart.Item{
						Sku:       "carvercandle-oxemerald-queenblaze-xl",
						Quantity:  1,
						Price:     11293,
						ExpiresAt: nil,
					},
				}
				ctx := context.Background()
				response, err := srv.AddItem(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				sessionId = response.SessionId
			})

			It("should get a cart from cache", func() {
				in := &cart.GetBySessionIdRequest{
					SessionId: sessionId,
					CartId:    cartId,
				}
				ctx := context.Background()
				response, err := srv.GetBySessionId(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(len(response.Cart.Items)).To(Equal(1))
				sessionId = response.SessionId
			})
		})

		Describe("AttachCustomer", func() {
			It("should attach a customer to cart", func() {
				in := &cart.AttachCustomerRequest{
					CartId:        cartId,
					CustomerEmail: "test@k8scommerce.com",
				}
				ctx := context.Background()
				response, err := srv.AttachCustomer(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				sessionId = response.SessionId
			})

			It("should not attach an invalid customer to cart", func() {
				in := &cart.AttachCustomerRequest{
					CartId:        cartId,
					CustomerEmail: "bob@k8scommerce.com",
				}
				ctx := context.Background()
				response, err := srv.AttachCustomer(ctx, in)
				Expect(err).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
		})

		Describe("BulkAddItemLogic", func() {
			It("should add an item to cart", func() {
				in := &cart.BulkAddItemsRequest{
					CartId: cartId,
					Items: []*cart.Item{
						{
							Sku:       "carvercandle-oxemerald-queenblaze-xl",
							Quantity:  1,
							Price:     11293,
							ExpiresAt: nil,
						},
						{
							Sku:       "samuraibald-scorpionsapphire-eaterboom",
							Quantity:  400,
							Price:     15887,
							ExpiresAt: nil,
						},
					},
				}
				ctx := context.Background()
				response, err := srv.BulkAddItems(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				sessionId = response.SessionId
			})

			It("should get a cart from cache", func() {
				in := &cart.GetBySessionIdRequest{
					SessionId: sessionId,
					CartId:    cartId,
				}
				ctx := context.Background()
				response, err := srv.GetBySessionId(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(len(response.Cart.Items)).To(Equal(2))
				sessionId = response.SessionId
			})
		})

		Describe("UpdateItemQuantity", func() {
			It("should update an item's quantity", func() {
				in := &cart.UpdateItemQuantityRequest{
					CartId:   cartId,
					Sku:      "carvercandle-oxemerald-queenblaze-xl",
					Quantity: 30,
				}
				ctx := context.Background()
				response, err := srv.UpdateItemQuantity(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				sessionId = response.SessionId
			})

			It("should get a cart from cache", func() {
				in := &cart.GetBySessionIdRequest{
					SessionId: sessionId,
					CartId:    cartId,
				}
				ctx := context.Background()
				response, err := srv.GetBySessionId(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(len(response.Cart.Items)).To(Equal(2))

				for _, item := range response.Cart.Items {
					if item.Sku == "carvercandle-oxemerald-queenblaze-xl" {
						Expect(item.Quantity).To(Equal(int32(30)))
					}
				}

				sessionId = response.SessionId
			})
		})

		Describe("UpdateStatus", func() {
			It("should update an cart's status", func() {
				in := &cart.UpdateStatusRequest{
					CartId: cartId,
					Status: cart.CartStatus_abandoned,
				}
				ctx := context.Background()
				response, err := srv.UpdateStatus(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				sessionId = response.SessionId
			})

			It("should get a cart from cache", func() {
				in := &cart.GetBySessionIdRequest{
					SessionId: sessionId,
					CartId:    cartId,
				}
				ctx := context.Background()
				response, err := srv.GetBySessionId(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(len(response.Cart.Items)).To(Equal(2))
				Expect(response.Cart.Status).To(Equal(cart.CartStatus_abandoned))
				sessionId = response.SessionId
			})
		})

		Describe("RemoveItemLogic", func() {
			It("should remove an item from cart", func() {
				in := &cart.RemoveItemRequest{
					CartId: cartId,
					Sku:    "carvercandle-oxemerald-queenblaze-xl",
				}
				ctx := context.Background()
				response, err := srv.RemoveItem(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				sessionId = response.SessionId
			})

			It("should get a cart from cache", func() {
				in := &cart.GetBySessionIdRequest{
					SessionId: sessionId,
					CartId:    cartId,
				}
				ctx := context.Background()
				response, err := srv.GetBySessionId(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(len(response.Cart.Items)).To(Equal(1))
				sessionId = response.SessionId
			})
		})

		XDescribe("ClearCart", func() {
			It("should clear a cart", func() {
				in := &cart.ClearCartRequest{
					CartId: cartId,
				}
				ctx := context.Background()
				response, err := srv.ClearCart(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				sessionId = response.SessionId
			})

			It("should not contain any items in cart", func() {
				in := &cart.GetBySessionIdRequest{
					CartId:    cartId,
					SessionId: sessionId,
				}
				ctx := context.Background()
				response, err := srv.GetBySessionId(ctx, in)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(len(response.Cart.Items)).To(Equal(0))
			})
		})
	})
})
