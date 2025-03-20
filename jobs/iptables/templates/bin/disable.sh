#!/bin/bash

<% p("iptables").each do |table, chains|
	chains.each do |chain, rules| %>

iptables -t "<%= table %>" -F "pfbr-custom-<%= chain %>"

	<% end %>
<% end %>
