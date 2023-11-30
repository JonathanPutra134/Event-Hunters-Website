// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Rating is an object representing the database table.
type Rating struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID     null.Int  `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	EventID    null.Int  `boil:"event_id" json:"event_id,omitempty" toml:"event_id" yaml:"event_id,omitempty"`
	Rating     null.Int  `boil:"rating" json:"rating,omitempty" toml:"rating" yaml:"rating,omitempty"`
	RatingDate null.Time `boil:"rating_date" json:"rating_date,omitempty" toml:"rating_date" yaml:"rating_date,omitempty"`

	R *ratingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ratingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RatingColumns = struct {
	ID         string
	UserID     string
	EventID    string
	Rating     string
	RatingDate string
}{
	ID:         "id",
	UserID:     "user_id",
	EventID:    "event_id",
	Rating:     "rating",
	RatingDate: "rating_date",
}

var RatingTableColumns = struct {
	ID         string
	UserID     string
	EventID    string
	Rating     string
	RatingDate string
}{
	ID:         "rating.id",
	UserID:     "rating.user_id",
	EventID:    "rating.event_id",
	Rating:     "rating.rating",
	RatingDate: "rating.rating_date",
}

// Generated where

var RatingWhere = struct {
	ID         whereHelperint
	UserID     whereHelpernull_Int
	EventID    whereHelpernull_Int
	Rating     whereHelpernull_Int
	RatingDate whereHelpernull_Time
}{
	ID:         whereHelperint{field: "\"rating\".\"id\""},
	UserID:     whereHelpernull_Int{field: "\"rating\".\"user_id\""},
	EventID:    whereHelpernull_Int{field: "\"rating\".\"event_id\""},
	Rating:     whereHelpernull_Int{field: "\"rating\".\"rating\""},
	RatingDate: whereHelpernull_Time{field: "\"rating\".\"rating_date\""},
}

// RatingRels is where relationship names are stored.
var RatingRels = struct {
	Event string
	User  string
}{
	Event: "Event",
	User:  "User",
}

// ratingR is where relationships are stored.
type ratingR struct {
	Event *Event `boil:"Event" json:"Event" toml:"Event" yaml:"Event"`
	User  *User  `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*ratingR) NewStruct() *ratingR {
	return &ratingR{}
}

func (r *ratingR) GetEvent() *Event {
	if r == nil {
		return nil
	}
	return r.Event
}

func (r *ratingR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// ratingL is where Load methods for each relationship are stored.
type ratingL struct{}

var (
	ratingAllColumns            = []string{"id", "user_id", "event_id", "rating", "rating_date"}
	ratingColumnsWithoutDefault = []string{}
	ratingColumnsWithDefault    = []string{"id", "user_id", "event_id", "rating", "rating_date"}
	ratingPrimaryKeyColumns     = []string{"id"}
	ratingGeneratedColumns      = []string{}
)

type (
	// RatingSlice is an alias for a slice of pointers to Rating.
	// This should almost always be used instead of []Rating.
	RatingSlice []*Rating
	// RatingHook is the signature for custom Rating hook methods
	RatingHook func(context.Context, boil.ContextExecutor, *Rating) error

	ratingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ratingType                 = reflect.TypeOf(&Rating{})
	ratingMapping              = queries.MakeStructMapping(ratingType)
	ratingPrimaryKeyMapping, _ = queries.BindMapping(ratingType, ratingMapping, ratingPrimaryKeyColumns)
	ratingInsertCacheMut       sync.RWMutex
	ratingInsertCache          = make(map[string]insertCache)
	ratingUpdateCacheMut       sync.RWMutex
	ratingUpdateCache          = make(map[string]updateCache)
	ratingUpsertCacheMut       sync.RWMutex
	ratingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ratingAfterSelectHooks []RatingHook

var ratingBeforeInsertHooks []RatingHook
var ratingAfterInsertHooks []RatingHook

var ratingBeforeUpdateHooks []RatingHook
var ratingAfterUpdateHooks []RatingHook

var ratingBeforeDeleteHooks []RatingHook
var ratingAfterDeleteHooks []RatingHook

var ratingBeforeUpsertHooks []RatingHook
var ratingAfterUpsertHooks []RatingHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Rating) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ratingAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Rating) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ratingBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Rating) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ratingAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Rating) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ratingBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Rating) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ratingAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Rating) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ratingBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Rating) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ratingAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Rating) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ratingBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Rating) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ratingAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRatingHook registers your hook function for all future operations.
func AddRatingHook(hookPoint boil.HookPoint, ratingHook RatingHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		ratingAfterSelectHooks = append(ratingAfterSelectHooks, ratingHook)
	case boil.BeforeInsertHook:
		ratingBeforeInsertHooks = append(ratingBeforeInsertHooks, ratingHook)
	case boil.AfterInsertHook:
		ratingAfterInsertHooks = append(ratingAfterInsertHooks, ratingHook)
	case boil.BeforeUpdateHook:
		ratingBeforeUpdateHooks = append(ratingBeforeUpdateHooks, ratingHook)
	case boil.AfterUpdateHook:
		ratingAfterUpdateHooks = append(ratingAfterUpdateHooks, ratingHook)
	case boil.BeforeDeleteHook:
		ratingBeforeDeleteHooks = append(ratingBeforeDeleteHooks, ratingHook)
	case boil.AfterDeleteHook:
		ratingAfterDeleteHooks = append(ratingAfterDeleteHooks, ratingHook)
	case boil.BeforeUpsertHook:
		ratingBeforeUpsertHooks = append(ratingBeforeUpsertHooks, ratingHook)
	case boil.AfterUpsertHook:
		ratingAfterUpsertHooks = append(ratingAfterUpsertHooks, ratingHook)
	}
}

// One returns a single rating record from the query.
func (q ratingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Rating, error) {
	o := &Rating{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for rating")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Rating records from the query.
func (q ratingQuery) All(ctx context.Context, exec boil.ContextExecutor) (RatingSlice, error) {
	var o []*Rating

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Rating slice")
	}

	if len(ratingAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Rating records in the query.
func (q ratingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count rating rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ratingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if rating exists")
	}

	return count > 0, nil
}

// Event pointed to by the foreign key.
func (o *Rating) Event(mods ...qm.QueryMod) eventQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.EventID),
	}

	queryMods = append(queryMods, mods...)

	return Events(queryMods...)
}

// User pointed to by the foreign key.
func (o *Rating) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadEvent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ratingL) LoadEvent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRating interface{}, mods queries.Applicator) error {
	var slice []*Rating
	var object *Rating

	if singular {
		var ok bool
		object, ok = maybeRating.(*Rating)
		if !ok {
			object = new(Rating)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeRating)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeRating))
			}
		}
	} else {
		s, ok := maybeRating.(*[]*Rating)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeRating)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeRating))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ratingR{}
		}
		if !queries.IsNil(object.EventID) {
			args = append(args, object.EventID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ratingR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.EventID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.EventID) {
				args = append(args, obj.EventID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`events`),
		qm.WhereIn(`events.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Event")
	}

	var resultSlice []*Event
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Event")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for events")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for events")
	}

	if len(eventAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Event = foreign
		if foreign.R == nil {
			foreign.R = &eventR{}
		}
		foreign.R.Ratings = append(foreign.R.Ratings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.EventID, foreign.ID) {
				local.R.Event = foreign
				if foreign.R == nil {
					foreign.R = &eventR{}
				}
				foreign.R.Ratings = append(foreign.R.Ratings, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ratingL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRating interface{}, mods queries.Applicator) error {
	var slice []*Rating
	var object *Rating

	if singular {
		var ok bool
		object, ok = maybeRating.(*Rating)
		if !ok {
			object = new(Rating)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeRating)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeRating))
			}
		}
	} else {
		s, ok := maybeRating.(*[]*Rating)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeRating)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeRating))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ratingR{}
		}
		if !queries.IsNil(object.UserID) {
			args = append(args, object.UserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ratingR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.UserID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.UserID) {
				args = append(args, obj.UserID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Ratings = append(foreign.R.Ratings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.ID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Ratings = append(foreign.R.Ratings, local)
				break
			}
		}
	}

	return nil
}

// SetEvent of the rating to the related item.
// Sets o.R.Event to related.
// Adds o to related.R.Ratings.
func (o *Rating) SetEvent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Event) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"rating\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"event_id"}),
		strmangle.WhereClause("\"", "\"", 2, ratingPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.EventID, related.ID)
	if o.R == nil {
		o.R = &ratingR{
			Event: related,
		}
	} else {
		o.R.Event = related
	}

	if related.R == nil {
		related.R = &eventR{
			Ratings: RatingSlice{o},
		}
	} else {
		related.R.Ratings = append(related.R.Ratings, o)
	}

	return nil
}

// RemoveEvent relationship.
// Sets o.R.Event to nil.
// Removes o from all passed in related items' relationships struct.
func (o *Rating) RemoveEvent(ctx context.Context, exec boil.ContextExecutor, related *Event) error {
	var err error

	queries.SetScanner(&o.EventID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("event_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Event = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Ratings {
		if queries.Equal(o.EventID, ri.EventID) {
			continue
		}

		ln := len(related.R.Ratings)
		if ln > 1 && i < ln-1 {
			related.R.Ratings[i] = related.R.Ratings[ln-1]
		}
		related.R.Ratings = related.R.Ratings[:ln-1]
		break
	}
	return nil
}

// SetUser of the rating to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Ratings.
func (o *Rating) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"rating\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, ratingPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserID, related.ID)
	if o.R == nil {
		o.R = &ratingR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Ratings: RatingSlice{o},
		}
	} else {
		related.R.Ratings = append(related.R.Ratings, o)
	}

	return nil
}

// RemoveUser relationship.
// Sets o.R.User to nil.
// Removes o from all passed in related items' relationships struct.
func (o *Rating) RemoveUser(ctx context.Context, exec boil.ContextExecutor, related *User) error {
	var err error

	queries.SetScanner(&o.UserID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("user_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.User = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Ratings {
		if queries.Equal(o.UserID, ri.UserID) {
			continue
		}

		ln := len(related.R.Ratings)
		if ln > 1 && i < ln-1 {
			related.R.Ratings[i] = related.R.Ratings[ln-1]
		}
		related.R.Ratings = related.R.Ratings[:ln-1]
		break
	}
	return nil
}

// Ratings retrieves all the records using an executor.
func Ratings(mods ...qm.QueryMod) ratingQuery {
	mods = append(mods, qm.From("\"rating\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"rating\".*"})
	}

	return ratingQuery{q}
}

// FindRating retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRating(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Rating, error) {
	ratingObj := &Rating{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"rating\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, ratingObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from rating")
	}

	if err = ratingObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ratingObj, err
	}

	return ratingObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Rating) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rating provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ratingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ratingInsertCacheMut.RLock()
	cache, cached := ratingInsertCache[key]
	ratingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ratingAllColumns,
			ratingColumnsWithDefault,
			ratingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ratingType, ratingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ratingType, ratingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"rating\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"rating\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into rating")
	}

	if !cached {
		ratingInsertCacheMut.Lock()
		ratingInsertCache[key] = cache
		ratingInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Rating.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Rating) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ratingUpdateCacheMut.RLock()
	cache, cached := ratingUpdateCache[key]
	ratingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ratingAllColumns,
			ratingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update rating, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"rating\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ratingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ratingType, ratingMapping, append(wl, ratingPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update rating row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for rating")
	}

	if !cached {
		ratingUpdateCacheMut.Lock()
		ratingUpdateCache[key] = cache
		ratingUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q ratingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for rating")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for rating")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RatingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ratingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"rating\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ratingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in rating slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all rating")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Rating) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rating provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ratingColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	ratingUpsertCacheMut.RLock()
	cache, cached := ratingUpsertCache[key]
	ratingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ratingAllColumns,
			ratingColumnsWithDefault,
			ratingColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			ratingAllColumns,
			ratingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert rating, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ratingPrimaryKeyColumns))
			copy(conflict, ratingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"rating\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(ratingType, ratingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ratingType, ratingMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert rating")
	}

	if !cached {
		ratingUpsertCacheMut.Lock()
		ratingUpsertCache[key] = cache
		ratingUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Rating record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Rating) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Rating provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ratingPrimaryKeyMapping)
	sql := "DELETE FROM \"rating\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from rating")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for rating")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ratingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ratingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rating")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for rating")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RatingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ratingBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ratingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"rating\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ratingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rating slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for rating")
	}

	if len(ratingAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Rating) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRating(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RatingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RatingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ratingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"rating\".* FROM \"rating\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ratingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RatingSlice")
	}

	*o = slice

	return nil
}

// RatingExists checks if the Rating row exists.
func RatingExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"rating\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if rating exists")
	}

	return exists, nil
}

// Exists checks if the Rating row exists.
func (o *Rating) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RatingExists(ctx, exec, o.ID)
}
