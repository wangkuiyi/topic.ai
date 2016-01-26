GOOS=linux GOARCH=amd64 go install github.com/wangkuiyi/topic.ai &&
    ssh root@topic.ai 'killall topic.ai' ;
    scp $GOPATH/bin/linux_amd64/topic.ai root@topic.ai:/root/ &&
    ssh root@topic.ai 'nohup /root/topic.ai'
