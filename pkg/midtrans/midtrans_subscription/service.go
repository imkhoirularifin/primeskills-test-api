package midtrans_subscription

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"primeskills-test-api/pkg/config"
	"primeskills-test-api/pkg/midtrans/midtrans_subscription/dto"
	"primeskills-test-api/pkg/utils"
	"primeskills-test-api/pkg/xlogger"

	"github.com/rs/zerolog"
)

type IMidtransSubscriptionService interface {
	CreateSubscription(req dto.CreateSubscriptionRequest) (*dto.CreateSubscriptionResponse, error)
	GetSubscriptionById(subscriptionId string) (*dto.GetSubscriptionResponse, error)
	DisableSubscription(subscriptionId string) (*dto.Response, error)
	CancelSubscription(subscriptionId string) (*dto.Response, error)
	EnableSubscription(subscriptionId string) (*dto.Response, error)
	UpdateSubscription(subscriptionId string, req dto.UpdateSubscriptionRequest) (*dto.Response, error)
}

var accessToken string

type service struct {
	log        *zerolog.Logger
	cfg        config.Config
	httpClient *http.Client
}

func (s service) CreateSubscription(req dto.CreateSubscriptionRequest) (*dto.CreateSubscriptionResponse, error) {
	url := fmt.Sprintf("%s/v1/subscriptions", s.cfg.Midtrans.BaseUrl)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", accessToken))

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res dto.CreateSubscriptionResponse
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create subscription: %s", *res.StatusMessage)
	}

	return &res, nil
}

func (s service) GetSubscriptionById(subscriptionId string) (*dto.GetSubscriptionResponse, error) {
	url := fmt.Sprintf("%s/v1/subscriptions/%s", s.cfg.Midtrans.BaseUrl, subscriptionId)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization",
		fmt.Sprintf("Basic %s", accessToken))

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res dto.GetSubscriptionResponse
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get subscription by id: %s", *res.StatusMessage)
	}

	return &res, nil
}

func (s service) DisableSubscription(subscriptionId string) (*dto.Response, error) {
	url := fmt.Sprintf("%s/v1/subscriptions/%s/disable", s.cfg.Midtrans.BaseUrl, subscriptionId)

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", accessToken))

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res dto.Response
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to disable subscription: %s", *res.StatusMessage)
	}

	return &res, nil
}

func (s service) CancelSubscription(subscriptionId string) (*dto.Response, error) {
	url := fmt.Sprintf("%s/v1/subscriptions/%s/cancel", s.cfg.Midtrans.BaseUrl, subscriptionId)

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", accessToken))

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res dto.Response
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to cancel subscription: %s", *res.StatusMessage)
	}

	return &res, nil
}

func (s service) EnableSubscription(subscriptionId string) (*dto.Response, error) {
	url := fmt.Sprintf("%s/v1/subscriptions/%s/enable", s.cfg.Midtrans.BaseUrl, subscriptionId)

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", accessToken))

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res dto.Response
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to enable subscription: %s", *res.StatusMessage)
	}

	return &res, nil
}

func (s service) UpdateSubscription(subscriptionId string, req dto.UpdateSubscriptionRequest) (*dto.Response, error) {
	url := fmt.Sprintf("%s/v1/subscriptions/%s", s.cfg.Midtrans.BaseUrl, subscriptionId)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization",
		fmt.Sprintf("Basic %s", accessToken))

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res dto.Response
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update subscription: %s", *res.StatusMessage)
	}

	return &res, nil
}

func NewService() IMidtransSubscriptionService {
	accessToken = utils.GenMidtransAccessToken(config.Cfg)
	httpClient := &http.Client{}

	return &service{
		cfg:        config.Cfg,
		log:        xlogger.Logger,
		httpClient: httpClient,
	}
}
