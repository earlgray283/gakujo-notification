package gakujo

import (
	"context"
	"os"

	"github.com/chromedp/chromedp"
)

type Client struct {
	id              string
	password        string
	ctx             context.Context
	allocCancelFunc context.CancelFunc
	taskCancelFunc  context.CancelFunc
}

func NewClient(ctx context.Context, id, password string) (*Client, error) {
	tempDirPath, err := os.MkdirTemp("", "gakujo-notification_")
	if err != nil {
		return nil, err
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.UserDataDir(tempDirPath),
		chromedp.Flag("headless", true),
	)
	allocCtx, allocCancelFunc := chromedp.NewExecAllocator(ctx, opts...)
	taskCtx, taskCancelFunc := chromedp.NewContext(
		allocCtx,
		// chromedp.WithErrorf(log.Printf),
		// chromedp.WithDebugf(log.Printf),
	)
	return &Client{
		ctx:             taskCtx,
		allocCancelFunc: allocCancelFunc,
		taskCancelFunc:  taskCancelFunc,
		id:              id,
		password:        password,
	}, nil
}

// must call this as a defer func
func (c *Client) Cancel() {
	c.allocCancelFunc()
	c.taskCancelFunc()
}
