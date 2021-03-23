#Principle of the Engage Digital App SDK

RingCentral's Engage Digital App SDK provides a way to customize RingCentral's Engage Digital UI behaviors by a third party developer using JavaScript and data APIs. Here, we will describe what is possible and how to do it and also what is not possible.

This Engage Digital App SDK provides a way to add buttons, tabs, open iframes, etc... in the Engage Digital (ED) agent interface. As a third party developer you will also be able to access the data via APIs. Everything will be written in JavaScript by the third party developer and executed in the agent context. This means that the RingCentral Engage Digital server does not need to access your application or data, only the agent, which makes it particularly suitable for internal IT (CRM ...) integration

Only a subset of JavaScript is allowed for your applications. All JavaScript methods and global objects are disabled except: ​`console`​ and ​`JSON`​. Therefore, you will **​not**​ be able to access or modify the `dom` of the page, reach the window object or make any AJAX requests. The purpose is to provide a safe sandboxed environment for our customers and avoid complex dependency on our HTML structure and/or unsafe interaction with external sources.

Please note that events are only issued by the browser, which means that events that occur on the server or elsewhere may not trigger the callbacks added through the SDK.
RingCentral's Engage Digital App SDK provides a list of high level JavaScript methods that are available on top of the standard JavaScript language. An exhaustive list of [Engage Digital APP Methods](./methods.md) can be found in the link. Those methods can be reached via the global S​MCC ​JavaScript object.

## Example
```
console.debug(SMCC.version())
```
