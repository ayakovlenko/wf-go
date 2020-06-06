// Common
// ------
var getTypeSignature = function (args) {
  var getType = function (x) {
    var t = Object.prototype.toString.call(x);
    return t.substring(8, t.length - 1);
  };
  args = Array.prototype.slice.apply(args);
  return args.map(getType).join(", ");
};

// Item
// ----

function __Item(args) {
  var signature = getTypeSignature(args);
  switch (signature) {
    case "String":
      break;
    case "String, String":
      this.note = args[1];
      break;
    case "String, Array":
      this.items = args[1];
      break;
    case "String, String, Array":
      this.note = args[1];
      this.items = args[2];
      break;
    default:
      throw new Error("unknown signature: Item(" + signature + ")");
  }
  this.title = args[0];

  // filter items
  if (this.items) {
    this.items = this.items.filter(function (item) {
      return Boolean(item);
    });
  }
}

function Item() {
  return new __Item(arguments);
}

__Item.prototype.add = function () {
  var item;
  var signature = getTypeSignature(arguments);
  switch (signature) {
    case "String":
      item = Item(arguments[0]);
      break;
    case "Object":
      item = arguments[0];
      break;
    default:
      throw new Error("unknown signature: Item.add(" + signature + ")");
  }

  if (this.items) {
    this.items.push(item);
    return;
  }

  this.items = [item];

  return this;
};

// Date
// ----

Date.prototype.isMonday = function () {
  return this.getDay() === MONDAY;
};

Date.prototype.isTuesday = function () {
  return this.getDay() === TUESDAY;
};

Date.prototype.isWednesday = function () {
  return this.getDay() === WEDNESDAY;
};

Date.prototype.isThursday = function () {
  return this.getDay() === THURSDAY;
};

Date.prototype.isFriday = function () {
  return this.getDay() === FRIDAY;
};

Date.prototype.isSaturday = function () {
  return this.getDay() === SATURDAY;
};

Date.prototype.isSunday = function () {
  return this.getDay() === SUNDAY;
};

Date.prototype.getDayName = (function () {
  var DAY_NAMES = [
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
  ];

  return function () {
    return DAY_NAMES[this.getDay()];
  };
})();
