# Principle of the RingCX Digital App SDK

The RingCX Digital App SDK allows third-party developers to customize RingCX Digital UI behavior using JavaScript and data APIs. This guide describes the available capabilities and their limitations.

The RingCX Digital App SDK can add buttons, tabs, and embedded frames to the agent interface. Third-party developers can also access data through APIs. The integration code runs in the agent's browser, which makes the SDK suitable for CRM and other internal application integrations.

Only a subset of JavaScript is allowed for your applications. All JavaScript methods and global objects are disabled except: ​`console`​ and ​`JSON`​. Therefore, you will **​not**​ be able to access or modify the `dom` of the page, reach the window object or make any AJAX requests. The purpose is to provide a safe sandboxed environment for our customers and avoid complex dependency on our HTML structure and/or unsafe interaction with external sources.

Please note that events are only issued by the browser, which means that events that occur on the server or elsewhere may not trigger the callbacks added through the SDK.
The RingCX Digital App SDK provides high-level JavaScript methods through the global `SMCC` object.

## Example
```
console.debug(SMCC.version())
```
