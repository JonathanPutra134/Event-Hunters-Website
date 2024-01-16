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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// EventCategory is an object representing the database table.
type EventCategory struct {
	ID         int `boil:"id" json:"id" toml:"id" yaml:"id"`
	CategoryID int `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	EventID    int `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`

	R *eventCategoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L eventCategoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EventCategoryColumns = struct {
	ID         string
	CategoryID string
	EventID    string
}{
	ID:         "id",
	CategoryID: "category_id",
	EventID:    "event_id",
}

var EventCategoryTableColumns = struct {
	ID         string
	CategoryID string
	EventID    string
}{
	ID:         "event_category.id",
	CategoryID: "event_category.category_id",
	EventID:    "event_category.event_id",
}

// Generated where

var EventCategoryWhere = struct {
	ID         whereHelperint
	CategoryID whereHelperint
	EventID    whereHelperint
}{
	ID:         whereHelperint{field: "\"event_category\".\"id\""},
	CategoryID: whereHelperint{field: "\"event_category\".\"category_id\""},
	EventID:    whereHelperint{field: "\"event_category\".\"event_id\""},
}

// EventCategoryRels is where relationship names are stored.
var EventCategoryRels = struct {
	Category string
	Event    string
}{
	Category: "Category",
	Event:    "Event",
}

// eventCategoryR is where relationships are stored.
type eventCategoryR struct {
	Category *Category `boil:"Category" json:"Category" toml:"Category" yaml:"Category"`
	Event    *Event    `boil:"Event" json:"Event" toml:"Event" yaml:"Event"`
}

// NewStruct creates a new relationship struct
func (*eventCategoryR) NewStruct() *eventCategoryR {
	return &eventCategoryR{}
}

func (r *eventCategoryR) GetCategory() *Category {
	if r == nil {
		return nil
	}
	return r.Category
}

func (r *eventCategoryR) GetEvent() *Event {
	if r == nil {
		return nil
	}
	return r.Event
}

// eventCategoryL is where Load methods for each relationship are stored.
type eventCategoryL struct{}

var (
	eventCategoryAllColumns            = []string{"id", "category_id", "event_id"}
	eventCategoryColumnsWithoutDefault = []string{"category_id", "event_id"}
	eventCategoryColumnsWithDefault    = []string{"id"}
	eventCategoryPrimaryKeyColumns     = []string{"id"}
	eventCategoryGeneratedColumns      = []string{}
)

type (
	// EventCategorySlice is an alias for a slice of pointers to EventCategory.
	// This should almost always be used instead of []EventCategory.
	EventCategorySlice []*EventCategory
	// EventCategoryHook is the signature for custom EventCategory hook methods
	EventCategoryHook func(context.Context, boil.ContextExecutor, *EventCategory) error

	eventCategoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	eventCategoryType                 = reflect.TypeOf(&EventCategory{})
	eventCategoryMapping              = queries.MakeStructMapping(eventCategoryType)
	eventCategoryPrimaryKeyMapping, _ = queries.BindMapping(eventCategoryType, eventCategoryMapping, eventCategoryPrimaryKeyColumns)
	eventCategoryInsertCacheMut       sync.RWMutex
	eventCategoryInsertCache          = make(map[string]insertCache)
	eventCategoryUpdateCacheMut       sync.RWMutex
	eventCategoryUpdateCache          = make(map[string]updateCache)
	eventCategoryUpsertCacheMut       sync.RWMutex
	eventCategoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var eventCategoryAfterSelectHooks []EventCategoryHook

var eventCategoryBeforeInsertHooks []EventCategoryHook
var eventCategoryAfterInsertHooks []EventCategoryHook

var eventCategoryBeforeUpdateHooks []EventCategoryHook
var eventCategoryAfterUpdateHooks []EventCategoryHook

var eventCategoryBeforeDeleteHooks []EventCategoryHook
var eventCategoryAfterDeleteHooks []EventCategoryHook

var eventCategoryBeforeUpsertHooks []EventCategoryHook
var eventCategoryAfterUpsertHooks []EventCategoryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EventCategory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventCategoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EventCategory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventCategoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EventCategory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventCategoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EventCategory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventCategoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EventCategory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventCategoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EventCategory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventCategoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EventCategory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventCategoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EventCategory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventCategoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EventCategory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventCategoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEventCategoryHook registers your hook function for all future operations.
func AddEventCategoryHook(hookPoint boil.HookPoint, eventCategoryHook EventCategoryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		eventCategoryAfterSelectHooks = append(eventCategoryAfterSelectHooks, eventCategoryHook)
	case boil.BeforeInsertHook:
		eventCategoryBeforeInsertHooks = append(eventCategoryBeforeInsertHooks, eventCategoryHook)
	case boil.AfterInsertHook:
		eventCategoryAfterInsertHooks = append(eventCategoryAfterInsertHooks, eventCategoryHook)
	case boil.BeforeUpdateHook:
		eventCategoryBeforeUpdateHooks = append(eventCategoryBeforeUpdateHooks, eventCategoryHook)
	case boil.AfterUpdateHook:
		eventCategoryAfterUpdateHooks = append(eventCategoryAfterUpdateHooks, eventCategoryHook)
	case boil.BeforeDeleteHook:
		eventCategoryBeforeDeleteHooks = append(eventCategoryBeforeDeleteHooks, eventCategoryHook)
	case boil.AfterDeleteHook:
		eventCategoryAfterDeleteHooks = append(eventCategoryAfterDeleteHooks, eventCategoryHook)
	case boil.BeforeUpsertHook:
		eventCategoryBeforeUpsertHooks = append(eventCategoryBeforeUpsertHooks, eventCategoryHook)
	case boil.AfterUpsertHook:
		eventCategoryAfterUpsertHooks = append(eventCategoryAfterUpsertHooks, eventCategoryHook)
	}
}

// One returns a single eventCategory record from the query.
func (q eventCategoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EventCategory, error) {
	o := &EventCategory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for event_category")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EventCategory records from the query.
func (q eventCategoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (EventCategorySlice, error) {
	var o []*EventCategory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EventCategory slice")
	}

	if len(eventCategoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EventCategory records in the query.
func (q eventCategoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count event_category rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q eventCategoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if event_category exists")
	}

	return count > 0, nil
}

// Category pointed to by the foreign key.
func (o *EventCategory) Category(mods ...qm.QueryMod) categoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CategoryID),
	}

	queryMods = append(queryMods, mods...)

	return Categories(queryMods...)
}

// Event pointed to by the foreign key.
func (o *EventCategory) Event(mods ...qm.QueryMod) eventQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.EventID),
	}

	queryMods = append(queryMods, mods...)

	return Events(queryMods...)
}

// LoadCategory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (eventCategoryL) LoadCategory(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEventCategory interface{}, mods queries.Applicator) error {
	var slice []*EventCategory
	var object *EventCategory

	if singular {
		var ok bool
		object, ok = maybeEventCategory.(*EventCategory)
		if !ok {
			object = new(EventCategory)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEventCategory)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEventCategory))
			}
		}
	} else {
		s, ok := maybeEventCategory.(*[]*EventCategory)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEventCategory)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEventCategory))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &eventCategoryR{}
		}
		args = append(args, object.CategoryID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &eventCategoryR{}
			}

			for _, a := range args {
				if a == obj.CategoryID {
					continue Outer
				}
			}

			args = append(args, obj.CategoryID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`category`),
		qm.WhereIn(`category.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Category")
	}

	var resultSlice []*Category
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Category")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for category")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for category")
	}

	if len(categoryAfterSelectHooks) != 0 {
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
		object.R.Category = foreign
		if foreign.R == nil {
			foreign.R = &categoryR{}
		}
		foreign.R.EventCategories = append(foreign.R.EventCategories, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CategoryID == foreign.ID {
				local.R.Category = foreign
				if foreign.R == nil {
					foreign.R = &categoryR{}
				}
				foreign.R.EventCategories = append(foreign.R.EventCategories, local)
				break
			}
		}
	}

	return nil
}

// LoadEvent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (eventCategoryL) LoadEvent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEventCategory interface{}, mods queries.Applicator) error {
	var slice []*EventCategory
	var object *EventCategory

	if singular {
		var ok bool
		object, ok = maybeEventCategory.(*EventCategory)
		if !ok {
			object = new(EventCategory)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEventCategory)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEventCategory))
			}
		}
	} else {
		s, ok := maybeEventCategory.(*[]*EventCategory)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEventCategory)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEventCategory))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &eventCategoryR{}
		}
		args = append(args, object.EventID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &eventCategoryR{}
			}

			for _, a := range args {
				if a == obj.EventID {
					continue Outer
				}
			}

			args = append(args, obj.EventID)

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
		foreign.R.EventCategories = append(foreign.R.EventCategories, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EventID == foreign.ID {
				local.R.Event = foreign
				if foreign.R == nil {
					foreign.R = &eventR{}
				}
				foreign.R.EventCategories = append(foreign.R.EventCategories, local)
				break
			}
		}
	}

	return nil
}

// SetCategory of the eventCategory to the related item.
// Sets o.R.Category to related.
// Adds o to related.R.EventCategories.
func (o *EventCategory) SetCategory(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Category) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"event_category\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"category_id"}),
		strmangle.WhereClause("\"", "\"", 2, eventCategoryPrimaryKeyColumns),
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

	o.CategoryID = related.ID
	if o.R == nil {
		o.R = &eventCategoryR{
			Category: related,
		}
	} else {
		o.R.Category = related
	}

	if related.R == nil {
		related.R = &categoryR{
			EventCategories: EventCategorySlice{o},
		}
	} else {
		related.R.EventCategories = append(related.R.EventCategories, o)
	}

	return nil
}

// SetEvent of the eventCategory to the related item.
// Sets o.R.Event to related.
// Adds o to related.R.EventCategories.
func (o *EventCategory) SetEvent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Event) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"event_category\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"event_id"}),
		strmangle.WhereClause("\"", "\"", 2, eventCategoryPrimaryKeyColumns),
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

	o.EventID = related.ID
	if o.R == nil {
		o.R = &eventCategoryR{
			Event: related,
		}
	} else {
		o.R.Event = related
	}

	if related.R == nil {
		related.R = &eventR{
			EventCategories: EventCategorySlice{o},
		}
	} else {
		related.R.EventCategories = append(related.R.EventCategories, o)
	}

	return nil
}

// EventCategories retrieves all the records using an executor.
func EventCategories(mods ...qm.QueryMod) eventCategoryQuery {
	mods = append(mods, qm.From("\"event_category\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"event_category\".*"})
	}

	return eventCategoryQuery{q}
}

// FindEventCategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEventCategory(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*EventCategory, error) {
	eventCategoryObj := &EventCategory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"event_category\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, eventCategoryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from event_category")
	}

	if err = eventCategoryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return eventCategoryObj, err
	}

	return eventCategoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EventCategory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no event_category provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventCategoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	eventCategoryInsertCacheMut.RLock()
	cache, cached := eventCategoryInsertCache[key]
	eventCategoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			eventCategoryAllColumns,
			eventCategoryColumnsWithDefault,
			eventCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(eventCategoryType, eventCategoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(eventCategoryType, eventCategoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"event_category\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"event_category\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into event_category")
	}

	if !cached {
		eventCategoryInsertCacheMut.Lock()
		eventCategoryInsertCache[key] = cache
		eventCategoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EventCategory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EventCategory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	eventCategoryUpdateCacheMut.RLock()
	cache, cached := eventCategoryUpdateCache[key]
	eventCategoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			eventCategoryAllColumns,
			eventCategoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update event_category, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"event_category\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, eventCategoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(eventCategoryType, eventCategoryMapping, append(wl, eventCategoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update event_category row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for event_category")
	}

	if !cached {
		eventCategoryUpdateCacheMut.Lock()
		eventCategoryUpdateCache[key] = cache
		eventCategoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q eventCategoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for event_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for event_category")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EventCategorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"event_category\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, eventCategoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in eventCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all eventCategory")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EventCategory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no event_category provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventCategoryColumnsWithDefault, o)

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

	eventCategoryUpsertCacheMut.RLock()
	cache, cached := eventCategoryUpsertCache[key]
	eventCategoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			eventCategoryAllColumns,
			eventCategoryColumnsWithDefault,
			eventCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			eventCategoryAllColumns,
			eventCategoryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert event_category, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(eventCategoryPrimaryKeyColumns))
			copy(conflict, eventCategoryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"event_category\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(eventCategoryType, eventCategoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(eventCategoryType, eventCategoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert event_category")
	}

	if !cached {
		eventCategoryUpsertCacheMut.Lock()
		eventCategoryUpsertCache[key] = cache
		eventCategoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EventCategory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EventCategory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EventCategory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), eventCategoryPrimaryKeyMapping)
	sql := "DELETE FROM \"event_category\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from event_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for event_category")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q eventCategoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no eventCategoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from event_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for event_category")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EventCategorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(eventCategoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"event_category\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventCategoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from eventCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for event_category")
	}

	if len(eventCategoryAfterDeleteHooks) != 0 {
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
func (o *EventCategory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEventCategory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventCategorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EventCategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"event_category\".* FROM \"event_category\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventCategoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EventCategorySlice")
	}

	*o = slice

	return nil
}

// EventCategoryExists checks if the EventCategory row exists.
func EventCategoryExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"event_category\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if event_category exists")
	}

	return exists, nil
}

// Exists checks if the EventCategory row exists.
func (o *EventCategory) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EventCategoryExists(ctx, exec, o.ID)
}