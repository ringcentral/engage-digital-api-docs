## Introduction to Locales

If you have implemented any custom actions you might be interested in providing translations for them. We have a simple mechanism to allow just that: add a `locales` attribute to the `implementation.info` response. This is completely optional and if you don't provide any translations we'll just [humanize](https://api.rubyonrails.org/classes/ActiveSupport/Inflector.html#method-i-humanize) the custom action names.

## Example Response with Translations:
```json
{
    "objects": {
        "messages": ["list", "show", "spam"],
        "threads": ["list", "show"]
    },
    "options": [],
    "locales": {
        "default": {"spam": "Mark as spam"},
        "fr": {"spam": "Marquer comme indésirable"}
    }
}
```

When we display a custom action `example` for a language `lang` we do the following:

* use `locales['lang']['example']` if defined
* use `locales['default']['example']` if defined
* use `example.humanize` otherwise.

Currently our interface is available in English and French.

Please bear in mind that this mechanism only works for custom actions.