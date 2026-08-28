# Sending Rich Text Messages

Agent messages sent via the [Create Content API](https://developers.ringcentral.com/engage/digital/api-reference/Contents/createContent) support an optional `body_format` parameter on **RingCX Web Messaging** and **RingCX Mobile Messaging** channels where rich text has been enabled. When set to `html`, the body is interpreted as HTML and rendered natively in the client.

## Request Example

```bash
curl -X POST "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents"
```

```json
{
  "in_reply_to_id": "<in_reply_to_id>",
  "body": "<strong>Hello</strong> <u>world</u>",
  "body_format": "html"
}
```

Omitting `body_format` (or setting it to `text`) treats the body as plain text, and the end customer will see the raw HTML instead of the rendered result.

!!! note
    Rich text must be enabled on the channel by an administrator before `body_format=html` can be used. Passing `body_format=html` on a channel without this feature returns a `422` error.

## Supported HTML

### Tags

| Tag | Purpose |
|-----|---------|
| `<p>`, `<br>` | Paragraphs and line breaks |
| `<strong>`, `<em>`, `<u>` | Bold, italic, underline |
| `<h1>` – `<h6>` | Headings |
| `<blockquote>` | Quoted text |
| `<ul>`, `<li>` | Unordered lists (`ol` is not supported) |
| `<a>` | Hyperlinks |
| `<span>` | Inline styling, such as colors. |
| `<hr>` | Horizontal rule |

Any tag not in this list is rejected with a `422` status and an explicit validation error.

### Attributes

| Attribute | Purpose |
|-----------|---------|
| `href` | Link target URL |
| `target` | Link opening behavior, e.g. `_blank` to open in new tab |
| `style` | To enrich an element with some style. Only `color` and `text-align` are accepted (see below) |

All other attributes are either stripped or produce a validation error.

### CSS properties

Only two CSS properties are allowed inside `style`:

| Property | Accepted values |
|----------|----------------|
| `color` | Six-digit hex only — `#RRGGBB` |
| `text-align` | `start`, `center`, `end` |

**Example of a valid HTML:**

```html
<p style="text-align:center">
  <span style="color:#E74C3C"><strong>Important notice</strong></span>
</p>
<ul>
  <li>First item</li>
  <li>Second item</li>
</ul>
<a href="https://example.com" target="_blank">Learn more</a>
```

## Validation errors

When the submitted HTML violates the allow-list, the API returns `422` with `error: validation_error` and an `errors` array describing each violation. Fix the offending markup and retry — no message is created until the body passes validation. Of course, when using `text`, the HTML inside the body is escaped and rendered as-is to the end-customer.
