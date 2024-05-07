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
