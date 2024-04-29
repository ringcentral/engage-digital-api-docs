# Surveys

RingCentral RingCX Messaging supports the integration with SurveyGizmo, like any other RingCentral RingCX source.

On **Messaging mode**, the survey is sent asynchronously as a message just like in other RingCentral RingCX sources.

On **Live Chat mode**, the survey is directly part of the visitor engagement workflow. This takes advantage of the real time capacity of this mode.

When a Survey is activated and eligible, it is presented to the visitor at the end of his conversation, inside the chat window, as soon as the agent closes the intervention.

As usual, the eligibility depends on the configuration (restriction by team, or categories). But contrary to asynchronous sendings, when eligible, a survey is always presented for an intervention (no anti spam feature), and only once per intervention (in case it is reopened).

Note that asynchronous sending mode is also possible, but its eligibility requires you to have the visitor email address.

!!! Abstract
    The goal was to go with our existing solution, SurveyGizmo, that is well used and highly customizable. Integrating it into the chat window allows direct feedback without distracting the visitor from the sale (important for the presales use case).

## Creating a Survey in SurveyGizmo

In SurveyGizmo, create a new survey and configure it with these specifics:

### Styling

In the “Style” main tab:

* In “Themes”, apply the “RingCentral RingCX Messaging - Base theme” that will serve as a base.
* In “Interactions”, choose “Standard” for Desktop Interaction and “Not mobile optimized” for Mobile Interaction.
* In other tabs, customize as wished

When the Survey will be created inside RingCentral RingCX, you’ll be able to preview the rendering for the exact sizing of the chat window.

### Building the questionnaire

In the “Build” main tab:

On the first page, add a text block to explain that the chat session is ending, inviting him to take a survey. Add this above the first question, or on its own dedicated first page. This is needed because the Survey is rendered in place of the “end of conversation” frame and the transition must be explained to avoid confusion.

Then, add a page per question, for a better flow.

On the last page, the “thank you” page, in addition to any “thank you” message, add a “JavaScript action” (in “Add an Action”), with description `RingCentral RingCX Messaging - Complete event`, with JavaScript action being:

``` javascript
if (window.parent) {
    window.parent.postMessage('dimelo:chat:survey:complete', '*');
}
```

The “thank you” page is shown for 3 seconds, then the survey frame is removed, letting the “end of conversation” frame be shown.

Choose one question to be used as the main satisfaction question (multiple choice options, with custom Reporting Values scaled from `1` to `5`; activate it in “Advanced Option Settings”) with its “Alias” set to `customer_satisfaction`. Refer to the survey documentation for further details.

## Adding the Survey in RingCentral RingCX

Add the Survey by its SurveyGizmo ID. Edit it right after to configure it by having a “Realtime” sending mode and selecting the RingCX Messaging source.

In the listing view, you’ll have a magnifier button to preview the Survey rendered with the exact sizing of the chat window. This is useful to adjust the customization and the questionnaire. Survey submitted in this preview mode will be tagged as “Test”.

## Multilingual survey

Survey, like RingCX Messaging, supports multi-language. The language of the survey is the language the chat is displayed in. It means:

* Chat **must** support this language (see [Custom Language](../source/#custom-language)), if you want to use a language that is not supported natively you must add it to the chat.
* Survey **must** be translated in any language the chat is supposed to support
