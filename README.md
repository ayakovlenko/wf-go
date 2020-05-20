# wf

_WorkFlowy CLI tool and library written in Go_

## Testing

```
go test ./... -coverprofile cp.out
```

https://gist.github.com/ayakovlenko/768b58f2d485c1224ee66e6b0a552107

## Templates

### Built-in functions

```ts
function dayName(day: int): string
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

Template parameter passed with `--param` / `-p` parameters can be accessed
through globally available `param` object:

```js
var date = new Date();

if (param.date === "tomorrow") {
  date.setDate(date.getDate() + 1);
}
```
