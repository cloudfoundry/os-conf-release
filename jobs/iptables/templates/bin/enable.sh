#!/bin/bash

set -e

function setup_chain {
	table=$1
	orig_chain=$2
	target_chain=$3

	if ! iptables -t "${table}" -L "${target_chain}" >/dev/null 2>&1; then
		iptables -t "${table}" -N "${target_chain}"
	fi

	if ! iptables -t "${table}" -C "${orig_chain}" -j "${target_chain}" 2>/dev/null; then
		iptables -t "${table}" -A "${orig_chain}" -j "${target_chain}"
	fi
}

<% p("iptables").each do |table, chains|
	chains.each do |chain, rules| %>

setup_chain "<%= table %>" "<%= chain %>" "pfbr-custom-<%= chain %>"

<% rules.each do |rule| %>
	iptables -t "${table}" -A "pfbr-custom-<%= chain %>"  <%= rule %>
<% end %>

	<% end %>
<% end %>
