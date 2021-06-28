# Creating Catalog Items

When the source is created, pre-built default items are created for reference and for an easy starting.

## Creating an Item - Common Part

You must provide a meaningful name in "Label" and choose the corresponding Engage Messaging source in "Source".

## Creating an Item: Invitation or a Button

### Available CSS Classes

* **`dimelo_chat_item_action`**: when set to an HTML element, it makes it the clickable action of the item (the element that launches the Chat). If this class is not set to any HTML element in the item, the whole item is clickable for action.
* **`dimelo_chat_item_dismiss`**: when set to an HTML element, it makes it the clickable dismiss action (the action to close and remove the item).

### Online / Offline Mode

When activated, the item HTML container has either the `dimelo_chat_mode_online` or `dimelo_chat_mode_offline` CSS classes you can use to style the item.

For example:

```css
.dimelo_chat_mode_online .my_button {
    background: green;
}
.dimelo_chat_mode_offline .my_button {
    background: red;
}
```

You have also the following utility CSS classes, to use directly in the HTML:

* **`dimelo_chat_show_when_online`**
* **`dimelo_chat_hide_when_online`**
* **`dimelo_chat_show_when_offline`**
* **`dimelo_chat_hide_when_offline`**