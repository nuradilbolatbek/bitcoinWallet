package handler

import (
	"bitcoinWallet"
	"bitcoinWallet/package/service"
	mock_service "bitcoinWallet/package/service/mocks"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
)

func TestHandler_createWallet(t *testing.T) {
	type mockBehavior func(r *mock_service.MockWalletManagement, userID int, wallet bitcoinWallet.Wallet)

	tests := []struct {
		name                 string
		userID               int
		inputBody            string
		inputWallet          bitcoinWallet.Wallet
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "CreateWallet_Success",
			userID:    1,
			inputBody: `{"balance": 231.421}`,
			inputWallet: bitcoinWallet.Wallet{
				Balance: 231.421,
			},
			mockBehavior: func(r *mock_service.MockWalletManagement, userID int, wallet bitcoinWallet.Wallet) {
				r.EXPECT().CreateWallet(userID, wallet).Return(1, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"id":1}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockWalletManagement(ctrl)
			test.mockBehavior(repo, test.userID, test.inputWallet)

			serv := &service.Service{WalletManagement: repo}

			handler := Handler{services: serv, mu: sync.Mutex{}}

			r := gin.New()
			r.POST("/wallets", func(c *gin.Context) {
				c.Set("userId", test.userID)

				handler.createWallet(c)
			})

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/wallets", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getWalletByUserId(t *testing.T) {
	type mockBehavior func(r *mock_service.MockWalletManagement, userID int, wallet bitcoinWallet.Wallet)

	tests := []struct {
		name                 string
		userID               int
		expectedBalance      float64
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:            "GetWalletByUserId_Success",
			userID:          1,
			expectedBalance: 231.421, // Assuming the expected balance for the user is 231.421
			mockBehavior: func(r *mock_service.MockWalletManagement, userID int, wallet bitcoinWallet.Wallet) {
				r.EXPECT().GetWalletByUserID(userID).Return(bitcoinWallet.Wallet{Balance: 231.421}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"balance":231.421}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockWalletManagement(ctrl)
			test.mockBehavior(repo, test.userID, bitcoinWallet.Wallet{})
			serv := &service.Service{WalletManagement: repo}

			handler := Handler{services: serv, mu: sync.Mutex{}}

			r := gin.New()
			r.GET("/wallets", func(c *gin.Context) {
				c.Set("userId", test.userID)

				handler.getWalletByUserId(c)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/wallets", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_depositToWallet(t *testing.T) {
	type mockBehavior func(r *mock_service.MockWalletManagement, walletID int, amount float64)

	tests := []struct {
		name                 string
		walletID             int
		inputAmount          float64
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "DepositToWallet_Success",
			walletID:    1,
			inputAmount: 50.0,
			mockBehavior: func(r *mock_service.MockWalletManagement, walletID int, amount float64) {
				r.EXPECT().DepositToWallet(walletID, amount).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":"ok"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockWalletManagement(ctrl)
			test.mockBehavior(repo, test.walletID, test.inputAmount)

			serv := &service.Service{WalletManagement: repo}

			handler := Handler{services: serv, mu: sync.Mutex{}}

			r := gin.New()
			r.PUT("/wallets/:id/deposit", func(c *gin.Context) {
				// Set wallet ID in context
				c.Set("id", strconv.Itoa(test.walletID))

				handler.depositToWallet(c)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/wallets/1/deposit", bytes.NewBufferString(fmt.Sprintf(`{"amount": %f}`, test.inputAmount)))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_withdrawFromWallet(t *testing.T) {
	type mockBehavior func(r *mock_service.MockWalletManagement, walletID int, amount float64)

	tests := []struct {
		name                 string
		walletID             int
		inputAmount          float64
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "DepositToWallet_Success",
			walletID:    1,
			inputAmount: 100.0,
			mockBehavior: func(r *mock_service.MockWalletManagement, walletID int, amount float64) {
				r.EXPECT().DepositToWallet(walletID, amount).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":"ok"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockWalletManagement(ctrl)
			test.mockBehavior(repo, test.walletID, test.inputAmount)

			serv := &service.Service{WalletManagement: repo}

			handler := Handler{services: serv, mu: sync.Mutex{}}

			r := gin.New()
			r.PUT("/wallets/:id/withdraw", func(c *gin.Context) {
				c.Set("id", strconv.Itoa(test.walletID))

				handler.depositToWallet(c)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/wallets/1/withdraw", bytes.NewBufferString(fmt.Sprintf(`{"amount": %f}`, test.inputAmount)))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
