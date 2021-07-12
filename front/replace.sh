#!/usr/bin/env sh

find '/app' -name '*.js' -exec sed -i -e 's,API_BASE_URL,'"$VUE_APP_ROOT_API"',g' {} \;
nginx -g "daemon off;"

find './dist' -name '*.js'