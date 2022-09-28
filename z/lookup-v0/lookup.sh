#!/bin/bash
# See which domains are on Cloudflare nameservers.

cat < z/testdata/domains.txt | \
    xargs -I{} dig NS {} | grep "IN\s*NS.*.cloudflare\.com" | \
    cut -f1 | sort | uniq