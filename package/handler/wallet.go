package handler

import (
	"bitcoinWallet"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createWallet(c *gin.Context) {

	h.mu.Lock()
	defer h.mu.Unlock()

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input bitcoinWallet.Wallet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.WalletManagement.CreateWallet(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getWalletByUserId(c *gin.Context) {
	h.mu.Lock()
	defer h.mu.Unlock()
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.services.WalletManagement.GetWalletByUserID(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"balance": lists.Balance,
	})

}

func (h *Handler) depositToWallet(c *gin.Context) {
	h.mu.Lock()
	defer h.mu.Unlock()

	walletId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid wallet ID")
		return
	}

	var input struct {
		Amount float64 `json:"amount"`
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Amount <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid amount")
		return
	}

	if err := h.services.WalletManagement.DepositToWallet(walletId, input.Amount); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
func (h *Handler) withdrawFromWallet(c *gin.Context) {
	h.mu.Lock()
	defer h.mu.Unlock()
	walletId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid wallet ID")
		return
	}

	var input struct {
		Amount float64 `json:"amount"`
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.services.WalletManagement.GetWalletByUserID(userId)

	if input.Amount > lists.Balance {
		newErrorResponse(c, http.StatusBadRequest, "not enough money in wallet")
		return
	}

	if input.Amount <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid amount")
		return
	}

	if err := h.services.WalletManagement.WithdrawFromWallet(walletId, input.Amount); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
