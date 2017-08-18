$ErrorActionPreference = "Stop";
trap { $host.SetShouldExit(1) }

$USERS=@()
<% p('users').each do |user| %>
  $USERS +="<%= user["name"] %>"
<% end %>

 [Reflection.Assembly]::LoadWithPartialName("System.Web")
 Foreach ($user in $USERS) {
   do {
       $pwd = [System.Web.Security.Membership]::GeneratePassword(10,2)
    } until ($pwd -match '\d')
   net user /add $user $pwd
   net localgroup administrators $user /add
 }
