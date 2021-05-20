# Usage

```scp ./file root@104.236.195.69/root/ts/data/source
ssh-add -K ~/.ssh/id_rsa
ssh -A root@104.236.195.69
cd /root/ts
./ts
exit
scp root@104.236.195.69/root/ts/data/result/sent/file_result ./file_result 
scp root@104.236.195.69/root/ts/data/result/report/file_result ./file_result
```
