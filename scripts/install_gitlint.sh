if [ ! -f "/usr/local/bin/gitlint" ]; then
  curl https://raw.githubusercontent.com/llorllale/go-gitlint/master/download-gitlint.sh > download-gitlint.sh && chmod +x download-gitlint.sh
  ./download-gitlint.sh -b /usr/local/bin
  rm download-gitlint.sh
fi


