先在本地做前三步，然后在github上创建新repo，再做第四第五步
1. git init
2. git add
3. git commit -m
4. git remote add origin git@xxx:yyy/repo_name.git
    - xxx: local authdication name/host (found in .ssh/config)
    - yyy: github account username
5. git push -u origin main