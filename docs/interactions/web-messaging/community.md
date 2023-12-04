# Configuring the Community

To add a RingCentral Engage Messaging to an existing RingCentral Engage instance, you need to create a RingCentral Engage Messaging community. No specific configuration is required.

## Extra Values Mapping

When needed, and with a proper integration on the target website, information about the visitor’s identity can be provided to the chat client. Alongside the predefined attributes, you may register extra values, at will. The `Extra value mapping` section of the Community configuration deals with the handling of these extra values.

The technicalities will be detailed in the [integration section](../integration/#identity-integration-optional), but, those extra values are given by
key-value pairs.

You may configure a mapping between some of the extra values registered on the identities and some of the predefined fields of an Identity Group shown to the agents in RingCentral Engage, or some of the configured Custom Fields. You may add as many mappings as needed by entering the key of the extra value and by selecting the destination field. Refer to the dedicated documentation for further details.

<img class="img-fluid" width="591" src="../../../img/web-messaging-extra-value-mapping.png">

### Map the UUID to an Identity Group Field

If you identify your customers through this ID, you can add your UUID to the extra values or directly map it to an Identity Group Field as if it was an extra value. If you use this method, it’s highly recommended to use a Custom Field with a Custom Regex matching your ID to avoid getting the Automated Dimelo UUID into your field.

<img class="img-fluid" width="627" src="../../../img/web-messaging-uuid-mapping.png">

## JWT Signature

RingCentral Engage Messaging supports registering information about the visitor’s identity with a JSON Web Token (JWT), a cryptographically secure signature mechanism. RingCentral EngageMessaging supports the HMAC SHA256 signature and the RSA SHA256 signature.

!!! warning
    When registering information about the visitor’s Identity, the data flow is:
    **the target website → the visitor’s browser → RingCentral Engage**
    You can’t have the guarantee that the data is not altered between the information that is announced on the target website and read by the Chat Client, and what is received by the Server. The only way to have such a guarantee is to add a cryptographically secure signature of the data. JSON Web Token (JWT) is such a signature mechanism, de facto standard, based on a shared secret or an asymmetric public/private key.

### Handling

The `Signed identities` field configures the handling of those signed identities by the Server:

* **Reject JWT signed identities** : only non-signed information is accepted. This is a strict mode. This is the default value.
* **Require JWT signed identities** : only signed information is accepted. This is a strict secure mode.

The JWT are required to have an expiry date set in an “exp” field AND the expiry date is required to be at least 1 hour in the future. Otherwise, the JWT will be rejected.

Since the JWT needs to expire, it is also better to renew it periodically by setting a new JWT. Otherwise the chat will display a message saying that the user needs to refresh the page to continue the conversation.

<img class="img-fluid" width="409" src="../../../img/web-messaging-session-expired.png">

### JWT Keys

JWT works with shared secrets, called `keys` , configurable in the `JWT keys` section:

<img class="img-fluid" width="476" src="../../../img/web-messaging-jwt-keys.png">

To add a JWT key, you have to set an identifier, the `Key ID`, and the key itself:

* Secret for HMAC. It must be a string of 24 characters (0-9 and a-f).
* Public key for RSA.

The `Key ID` must be a string containing no spaces (a-z 0-9 and ”_”).

You may add several keys. This is useful because keys are secret, must be compartmentalised to the utmost, and may change periodically.

Changing a key used by the target website without requiring synchronisation between the deployment of the new key on the target website and the changement in RingCentral Engage can be done by simply adding the new key in the configuration. The previous one can be removed at a later time, when it is sure that it is no longer in use.
