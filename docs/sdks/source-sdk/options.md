# Introduction to Options

There are some options avaialble for your unique integration. By default, all of the options are disabled.

## messages.unthreaded

Useful when your messages don't have a rigid structure (Twitter for example).

Message objects will not have to specify a thread (and it will be ignored if specified). This option will make us (among other things) ignore all `thread.*` actions, even if you specify them as implemented.

## messages.no_title

Forums, for example, allow you to specify a Title for your messages, while Wordpress comments don't have one. Having this option will allow us to show or hide the `title` field for agents replying.

## private_messages.no_title

Same as above, but for private messages.

## threads.no_content

Some applications are threaded, but their threads are only there to provide structure, without having extra data. Forums are an example of this, while blogs are the opposite - the thread (post) contains content and comments are messages sent as a reply.

What will change:

* the only fields parsed for thread objects will be `id` and `title`.
* only the `threads.show` and `threads.list` actions will be valid for threads, all others will be ignored.

## messages.text

Allow messages body to be text formatted.

!!! info "Note"
    You can specify multiple formats at a time, like this:
    ```json
    {
        ...
        "options": ["messages.text", "messages.html"],
        ...
    }
    ```

## messages.html

Allow message body to be html formatted.

This is the **default** format when format is not specified in the options.

!!! info "Note"
    JS and CSS are sanitized.

## messages.allow_survey

Allows sending a survey via [Messages](../objects/#messages). Can be used with only messages enabled.

## private_messages.text

Allow private message body to be text formatted.

## private_messages.html

Allow private message body to be html formatted.

This is the **default** format when format is not specified in the options.

!!! info "Note"
    JS and CSS are sanitized.

## threads.text

Allow thread body to be text formatted.

## threads.html

Allow thread body to be html formatted.

This is the **default** format when format is not specified in the options.

!!! info "Note"
    JS and CSS are sanitized.

## view.messaging

Enable the messaging view in Engage Digital. Can be used with **only with one resource enabled** (messages or private_messages) and with **text only**. Must be implemented if structured contents support is required.

## messaging.typing_indicator

Adds the ability to send/receive `typing` events (`typing.start`/`typing.stop`). **Can only be used if the [`view.messaging`](#viewmessaging) options is enabled**.

## messaging.queuing_update

Adds the ability to receive `messaging.queuing_update` requests (only available if using the Task mode in Engage Digital).


## messages.attachments

Enable the message composition with attachments in Engage Digital.

## private_messages.attachments

Enable the private message composition with attachments in Engage Digital.

## threads.attachments

Enable the thread composition with attachments in Engage Digital.
