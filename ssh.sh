ssh-keygen -t rsa -C 'javaer-in-go@outlook.com' -f ~/.ssh/id_rsa_javaeringo

eval `ssh-agent`
ssh-add ~/.ssh/id_rsa
ssh-add ~/.ssh/id_rsa_javaeringo