package ggmanager

import (
	"context"
	"fmt"
)

type GGManagerInterface interface {
	Run(ctx context.Context) error
}

type GGManager struct {
	config *Config
}

func Init() (*GGManager, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("Init: %w", err)
	}

	return &GGManager{
		config: config,
	}, nil
}

func (ggm *GGManager) Run(ctx context.Context) error {
	fmt.Println("Hellow World")
	return nil
}
