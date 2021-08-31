# Customization

We designed the chat window to respect industry standards allowing for a user to be comfortable and to ease the engagement.

In a first approach, we encourage a light customization, purely in CSS. We offer a mechanism to change the colors of the chat window, based on our default stylesheet. This is done directly in RingCentral Engage.

## Customizing the Languages

You can add your own languages or modify the texts of the current languages with the custom languages field in your Engage Messaging source or by implementing it in the JavaScript of the page containing the chat.

``` html
<script>
    var _chatq = _chatq || [];
    _chatq.push(["_addTranslations", {“en”:{ … });
</script>
```

You can see [a list](https://docs.google.com/document/d/1ly9bR9q7VSBOx4PdeUpp9HSt5PyxFW3r_T_lPdJOmgc/edit) of all translations on many languages.

Note: The JavaScript has priority over the configuration in the source.

On the next page is an example of JSON language format:

``` json
{
    "en": {
        "attachment_frame": {
            "drop_text": "Drop your file here"
        },
        "attachment_error": {
            "file_too_large": "File is too large (maximum %{max_size_mb} MB)",
            "file_type_not_allowed": "File type not allowed",
            "message": "Error while sending the file %{filename} :",
            "unknown": "Unknown error"
        },
        "cobrowsing": {
            "request_message": {
                "body": "Do you accept to share your screen?",
                "accept": "Yes",
                "deny": "No"
            }
        },
        "connection": "%{screenname} is now connected",
        "content": {
            "anonymized": "This content was anonymized upon author request."
        },
        "control_attachment_title": "Send a file",
        "control_back_title": "Conversations list",
        "control_close_cobrowsing": "Deactivate",
        "control_close_title": "Minimize",
        "control_emoji_title": "Add an emoji",
        "control_new_title": "New conversation",
        "control_quit_title": "Quit",
        "control_send_title": "Send",
        "conversation_header_title": "Your conversation",
        "date": {
            "ago_format": {
                "about_x_minutes": {
                    "one": "about 1 minute ago",
                    "other": "about %{count} minutes ago"
                },
                "default": "%m %d %Y %H:%M",
                "different_day": "%m %d",
                "different_year": "%m %Y",
                "now": "just now",
                "same_day": "Today %H:%M",
                "same_year": "%m %d %H:%M"
            },
            "estimated_wait_time_format": {
                "hours": {
                    "one": "1 hour",
                    "other": "more than one hour"
                },
                "minutes": {
                    "one": "1 minute",
                    "other": "%{count} minutes"
                }
            },
            "month_names": ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]
        },
        "default_cobrowsing_active_status": "Cobrowsing is active",
        "default_title": "Engage Messaging",
        "dialog_frame": {
            "button_ok_label": "Ok",
            "error_title": "Error !"
        },
        "disconnection": "%{screenname} is now disconnected",
        "end_frame": {
            "confirm": "Do you want to end the chat session ?",
            "message_1": "Thank you for using the chat.",
            "message_2": "The session has ended, here are your options:",
            "option_leave_label": "Leave the conversation",
            "option_print_label": "Print the conversation",
            "option_see_again_label": "See the conversation"
        },
        "input_placeholder": "Write your message here",
        "printer": {
            "title": "History of your conversation on %{title}"
        },
        "prompt_error": {
            "incorrect_email": "This email is incorrect",
            "required": "This field is required",
            "send_error": "An error occurred, please try again"
        },
        "queuing_update": {
            "estimated_wait_time": "Your estimated wait time is %{estimated_wait_time}",
            "queued": "You are in position %{position} in the waiting queue",
            "transfered": "You are being transfered"
        },
        "request_email": {
            "title": "Please enter your email address so that we can reach you:",
            "submitted_title": "Thank you for submitting your email. You will be notified by email."
        },
        "send_message_error": "An error happened, click to try again.",
        "thread_closure": "The conversation has ended.",
        "threads_listing_empty": "No conversations",
        "timed_messages": {
            "impossible": "Sorry, we were not able to find any available agent."
        },
        "typing": "%{screenname} is typing",
        "unsupported_browser": "Internet Explorer 9 and below are not supported",
        "waiting_for_agent": "Waiting for an agent",
        "you": "You",
        "you_colon": "You: "
    }
}
```

## Customizing the CSS

We extensively use SCSS that is a superset of the CSS language and that provides some useful features such as variables. We generate from SCSS the CSS used by the chat system.

On the Engage Messaging source configuration, there is a field named custom CSS. This field accepts CSS code as well as SCSS code. The code inserted here is used to generate a custom stylesheet for the chat window and to generate as well the stylesheets of each of the Catalog items (Button, Invitation…) you may use.

In this field, you can have some global CSS code and you can redefine the variables we use to generate the “design” part of the chat stylesheet, in order to recreate it with the colors wanted.

## Changing Colors of the Chat Window

### Details of Color Variables

Here is the full list of color related variables we use to generate the chat window stylesheet:

|**Variable**|**Default**|**Details**|
| - | - | - |
|$gray900|#212121||
|$gray800|#757575||
|$gray700|#A1A1A1||
|$gray400|#D1D1D1||
|$gray300|#E0E0E0||
|$gray200|#E7E7E7||
|$gray100|#EFEFF0||
|$gray50|#F6F7F8||
|$higlight-color|#4481eb|Main color|
|$background-color|#ffffff|Background of the Chat window|
|$highlight-text-color|#ffffff|Color for text displayed above a `$highlight-color` background|
|$text-color|$gray900|Main text color|
|$secondary-text-color|$gray700|Secondary text color (screen name, timestamp,…)|
|$placeholder-color|$gray700|Color of placeholder text in form fields.|
|$error-color|#f44336|Color for errors related elements.|
|$global-notification-color|#f0512a|Background color of the new message notification badges.|
|$end-conversation-option-color|$highlight-color|Color for the "See conversation" and "Print conversation" buttons|
|$leave-conversation-color|#f0512a|Color of the "leave conversation" button on livechat|
|$status-color|mix($highlight-color, $gray100, 20%)|Color for the status box|
|$thread-closure-color|$gray800|Color of the "end of conversation" text|
|$send-icon-color|#ffffff|Color of the send icon|
|$bubble-icon-color|#ffffff|Color of the bubble icons|
|$header-color|$highlight-color|Background color of the header|
|$header-text-color|$highlight-text-color|Text color of the header|
|$header-icon-color|#ffffff|Color of the header’s icons|
|$message-background-color|#ecf0f6|Background color of agent messages|
|$message-text-color|$text-color|Text color of agent messages|
|$message-self-background-color|$highlight-color|Background color of visitor messages|
|$message-self-text-color|$highlight-text-color|Text color of visitor messages|
|$cobrowsing-status-color|#667080|Color of the cobrowsing status.|
|$cobrowsing-status-button-bg-color|$gray100|Background color of the cobrowsing status|
|$webview-background-color|$background-color| Background color of the webview and its header|
|$webview-header-text-color|$text-color|Color of the webview header text|

You only have to define the one you want to change.

### Common Use Case

The common use case is to only change the `$highlight-color` as the others colors are pretty neutral and industry standard.

Other changes are beyond our recommendation.

## Changing Colors of the Default Catalog Items

### Details of Color Variables

|**Variable**|**Default**|**Details**|
| - | - | - |
|$highlight-color|#53cca4|Main color, used for buttons and interactions.|
