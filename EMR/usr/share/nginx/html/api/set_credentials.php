<?php
$ignoreAuth = true;
include_once("../interface/globals.php");
include_once("../library/authentication/password_hashing.php");
include_once("../library/authentication/password_change.php");
include_once("../library/user.inc");
include_once("../library/acl.inc");

require_once("$srcdir/sql.inc");

$json_data = json_decode(stripslashes(file_get_contents("php://input")), true);

$username = $json_data['username'];
$passwd = $json_data['password'];

$insertUserSQL=
    "insert into users set " .
    "username = '"         . $username .
    "', password = '"      . 'NoLongerUsed'                  .
    "', fname = '"         . 'user' .
    "', mname = '"         . 'user' .
    "', lname = '"         . 'user'.
    "', federaltaxid = '"  . '000-00-0000' .
    "', state_license_number = '"  . ''.
    "', newcrop_user_role = '"  . 'admin'.
    "', physician_type = '"  . '' .
    "', authorized = '"    . 0 .
    "', info = '"          . '' .
    "', federaldrugid = '" . ''.
    "', upin = '"          . '' .
    "', npi  = '"          . ''.
    "', taxonomy = '"      . '207Q00000X'.
    "', facility_id = '"   . 1 .
    "', specialty = '"     . ''.
    "', see_auth = '"      . ''.
    "', cal_ui = '"        . 3 .
    "', default_warehouse = '" . '' .
    "', irnpool = '"       . ''.
    "', calendar = '"      . 3   .
    "', pwd_expiration_date = '" . '0000-00-00'.
    "', import = TRUE";
$prov_id = "";

$success = update_password(1, 0, $passwd, $passwd, $passwd, true, $insertUserSQL, $username, $prov_id);
if($success)
{
    //set the facility name from the selected facility_id
    sqlStatement("UPDATE users, facility SET users.facility = facility.name WHERE facility.id = '" . 1 . "' AND users.username = '" . $username . "'");

    sqlStatement("insert into groups set name = '" . 'Default'.
        "', user = '" . $username . "'");

    add_user_aros($username, 'Administrators');
    if (isset($phpgacl_location)) {
        // Set the access control group of user
        set_user_aro('Default', $username,
            'user', 'user', 'user');


    }
}

