#!/bin/bash
# See which domains are on Cloudflare nameservers.

xargs -I{} dig NS {} | grep "IN\s*NS.*\.cloudflare\.com\."