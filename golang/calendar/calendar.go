package calendar

import (
    "time"
)

type Entry struct {
    start   time.Time
    end     time.Time
    title   string
    desc    string
    tz      string
}

func NewEntry(start, end time.Time, title, desc, tz string) *Entry {
    return &Entry{
        start: start,
        end: end,
        title: title,
        desc: desc,
        tz: tz,
    }
}

func (e *Entry) Start() time.Time {
    return e.start
}

func (e *Entry) End() time.Time {
    return e.end
}

func (e *Entry) Title() string {
    return e.title
}

func (e *Entry) Description() string {
    return e.desc
}

func (e *Entry) Timezone() string {
    return e.tz
}

type Calendar struct {
    name    string
    entries []Entry
}

func NewCalendar(name string) *Calendar {
    return &Calendar{
        name: name,
        entries: []Entry{},
    }
}

func (c *Calendar) Name() string {
    return c.name
}

func (c *Calendar) Entries() []Entry {
    return c.entries
}
