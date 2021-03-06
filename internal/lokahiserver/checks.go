package lokahiserver

import (
	"context"
	"errors"

	"github.com/Xe/lokahi/internal/database"
	"github.com/Xe/lokahi/rpc/lokahi"
	"github.com/Xe/uuid"
	"github.com/jinzhu/gorm"
)

var errNotImpl = errors.New("not implemented")

// Checks implements service Checks.
type Checks struct {
	DB *gorm.DB
}

// Create creates a new health check with the given options.
func (c *Checks) Create(ctx context.Context, opts *lokahi.CreateOpts) (*lokahi.Check, error) {
	dck := database.Check{
		UUID:        uuid.New(),
		URL:         opts.Url,
		WebhookURL:  opts.WebhookUrl,
		Every:       int(opts.Every),
		PlaybookURL: opts.PlaybookUrl,
		State:       lokahi.Check_INIT.String(),
	}

	err := c.DB.Create(&dck).Error
	if err != nil {
		return nil, err
	}

	return dck.AsProto(), nil
}

// Delete removes a check by ID and returns the data that was deleted.
func (c *Checks) Delete(ctx context.Context, cid *lokahi.CheckID) (*lokahi.Check, error) {
	dck, err := c.getCheck(ctx, cid.Id)
	if err != nil {
		return nil, err
	}

	err = c.DB.Unscoped().Where("uuid = ?", dck.UUID).Delete(database.Check{}).Error
	if err != nil {
		return nil, err
	}

	return dck.AsProto(), nil
}

// Get returns information on a check by ID.
func (c *Checks) Get(ctx context.Context, cid *lokahi.CheckID) (*lokahi.Check, error) {
	dck, err := c.getCheck(ctx, cid.Id)
	if err != nil {
		return nil, err
	}

	return dck.AsProto(), nil
}

// getCheck gets a check from the database
func (c *Checks) getCheck(ctx context.Context, id string) (*database.Check, error) {
	var ck database.Check
	err := c.DB.Where("uuid = ?", id).First(&ck).Error
	if err != nil {
		return nil, err
	}

	return &ck, nil
}

// List returns a page of checks based on a few options.
func (c *Checks) List(ctx context.Context, opts *lokahi.ListOpts) (*lokahi.ChecksPage, error) {
	var dchecks []database.Check

	err := c.DB.
		Limit(int(opts.Count)).
		Offset(int(opts.Page * opts.Count)).Find(&dchecks).Error
	if err != nil {
		return nil, err
	}

	result := &lokahi.ChecksPage{}

	for _, dc := range dchecks {
		result.Results = append(result.Results, &lokahi.ChecksPage_Result{Check: dc.AsProto()})
	}

	return result, nil
}

// Put updates a Check.
func (c *Checks) Put(ctx context.Context, chk *lokahi.Check) (*lokahi.Check, error) {
	dck, err := c.getCheck(ctx, chk.Id)
	if err != nil {
		return nil, err
	}

	dck.URL = chk.Url
	dck.WebhookURL = chk.WebhookUrl
	dck.WebhookResponseTimeNanoseconds = chk.WebhookResponseTimeNanoseconds
	dck.Every = int(chk.Every)
	dck.PlaybookURL = chk.PlaybookUrl
	dck.State = chk.State.String()

	err = c.DB.Save(dck).Error
	if err != nil {
		return nil, err
	}

	return dck.AsProto(), nil
}

// Status returns the detailed histogram status of a check.
func (c *Checks) Status(ctx context.Context, cid *lokahi.CheckID) (*lokahi.CheckStatus, error) {
	return nil, errNotImpl
}
