# wf

_organize your brain… as code_

`wf` can generate [WorkFlowy](https://workflowy.com) templates with arbitrary
logic.

By convention, templates must reside in `$WF_DIR/templates`, and the only
requirement is that they must return `Item` object.

The simplest possible template is:

```js
Item("Hello, World!");
```

You can use any valid JS as long as it's ES 5.1.
Under the hood, `wf` relies on [dop251/goja][goja] package which supports only
ES 5.1 yet.

```js
// $WF_DIR/templates/daily.js
var date = new Date();

var isoDate = date.toISOString().split("T")[0];

var todo = Item("To do today");

if (date.isSaturday()) {
  todo.add(Item("Laundry"));
}

if (date.isMonday() || date.isWednesday() || date.isFriday()) {
  todo.add(
    Item("Workout", [
      Item("Push-ups"),
      Item("Squats"),
      Item("Plank"),
    ])
  );
}

Item(
  isoDate,
  date.getDayName(),
  [
    Item("Menu"),
    todo,
  ]
);
```

To use the template, run `wf template` command giving a name of the template
file without `.js` extension:

```
$ wf template daily
```

The output of the command will be:

```xml
<opml version="1.0">
  <body>
    <outline text="2020-05-15" _note="Saturday">
      <outline text="Menu">
        <outline text="Meat"></outline>
      </outline>
      <outline text="To do today">
        <outline text="Laundry"></outline>
      </outline>
    </outline>
  </body>
</opml>
```

When you paste this XML is into WorkFlowy, it's going to be transformed into
a nice-looking list:

![](https://i.imgur.com/kTOwuIr.png)

## Templates

### Item DSL

You have 4 overloaded functions to define items:

```ts
function Item(title: string): Item

function Item(title: string, note: string): Item

function Item(title: string, items: Array<Item>): Item

function Item(title: string, note: string, items: Array<Item>): Item
```

### Built-in functions

```js
Date.prototype.getDayName

Date.prototype.isMonday

Date.prototype.isTuesday

Date.prototype.isWednesday

Date.prototype.isThursday

Date.prototype.isFriday

Date.prototype.isSaturday

Date.prototype.isSunday

Date.prototype.getDayName
```

### Built-in constants

```js
const MONDAY    = 1;
const TUESDAY   = 2;
const WEDNESDAY = 3;
const THURSDAY  = 4;
const FRIDAY    = 5;
const SATURDAY  = 6;
const SUNDAY    = 7;
```

### Parameters

Template parameters are passed after template name and have the following
syntax:

```bash
$ wf template <name> [key=value ...]
```

Example:

```bash
$ wf template daily date=tomorrow
```

To access parameters in templates, use global `param` object:

```js
var date = new Date();

if (param.date === "tomorrow") {
  date.setDate(date.getDate() + 1);
}
```

[goja]: https://github.com/dop251/goja
