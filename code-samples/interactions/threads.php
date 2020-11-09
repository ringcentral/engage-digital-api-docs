<?php
$SERVER = "https://<YOUR-DOMAIN>.api.engagement.dimelo.com";
$ACCESS_TOKEN = '<API-ACCESS-TOKEN>';
$API = "/1.0/content_threads";

try {
    $url = $SERVER . $API;
    $headers = array (
          'Authorization: Bearer ' . $ACCESS_TOKEN
        );
    try {
        $ch = curl_init($url);
        curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($ch, CURLOPT_HTTPGET, TRUE);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
        curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
        curl_setopt($ch, CURLOPT_TIMEOUT, 600);

        $strResponse = curl_exec($ch);
        $curlErrno = curl_errno($ch);
        if ($curlErrno) {
            throw new Exception($ecurlError);
        } else {
            $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
            curl_close($ch);
            if ($httpCode == 200) {
              print_r($strResponse."\n");
            }else{
              print_r($httpCode."\n");
            }
        }
    } catch (Exception $e) {
        throw $e;
    }
}catch (Exception $e) {
    throw $e;
}
?>