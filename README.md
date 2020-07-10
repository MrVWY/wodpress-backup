# wodpress-backup
 
#前言
+ 保证备份脚本在wp-config.php运行时所在目录下
+ 了解[wp-cli命令集](https://github.com/wp-cli/wp-cli)

#步骤
+ 编写备份脚本
  wp export --dir=your_Generate_address --user=your_name --post_type=post --filename_format=your_backupFileName.xml --allow-root
+ 配置scp传输免登录
+ 通过scp传输文件, 在脚本文件中添加
  + scp your_file root@IP:save_path
+ 设置定时任务
  + crontab -e 进入编辑模式
  + 0   12   *   *   *   /bin/sh  your .sh file path  
    分  时   日   月  周  |《===命令行====》|
  + ctrl+x Y 保存