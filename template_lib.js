function __Item(args) {
  var getTypeSignature = function (args) {
    var getType = function (x) {
      var t = Object.prototype.toString.call(x);
      return t.substring(8, t.length - 1);
    };
    args = Array.prototype.slice.apply(args);
    return args.map(getType).join(", ");
  };

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

__Item.prototype.add = function (item) {
  if (this.items) {
    this.items.push(item);
    return;
  }

  this.items = [item];
};

function Item() {
  return new __Item(arguments);
}
