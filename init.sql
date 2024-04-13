use S_T;/*将 S_T 设为当前数据库*/

create table Student 
 (Sno CHAR(9) PRIMARY KEY,
 Sname CHAR(20) UNIQUE,
 Ssex CHAR(2),
 Sage SMALLINT,
 Sdept CHAR(20),
 Scholarship char(2)
 )DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
 create table Course 
 (Cno CHAR(4) PRIMARY KEY,
 Cname CHAR(40),
 Cpno CHAR(4),
 Ccredit SMALLINT,
 FOREIGN KEY (Cpno) REFERENCES Course(Cno)
 )DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
 create table SC 
 (Sno CHAR(9),
 Cno CHAR(4),
 Grade SMALLINT,
 primary key (Sno, Cno),
 FOREIGN KEY (Sno) REFERENCES Student(Sno),
 FOREIGN KEY (Cno) REFERENCES Course(Cno)
 )DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;;

insert into Student values('200215121','李勇','男',20,'CS','否');
insert into Student values('200215122','刘晨','女',19,'CS','否');
insert into Student values('200215123','王敏','女',18,'MA','否');
insert into Student values('200215125','张立','男',19,'IS','否');

/*为表 Student 添加数据*/
insert into course values('1', '数据库', NULL,4);
insert into course values('2', '数学', NULL,2);

insert into Course values('3','信息系统',NULL,4);
insert into Course values('4','操作系统',NULL,3);
insert into Course values('5','数据结构',NULL,4);
insert into Course values('6','数据处理',NULL,2);
insert into Course values('7','java',NULL,4);

update Course set Cpno = '5' where Cno = '1';
update Course set Cpno = '1' where Cno = '3';
update Course set Cpno = '6' where Cno = '4';
update Course set Cpno = '7' where Cno = '5';
update Course set Cpno = '6' where Cno = '7';
/*为表 Course 添加数据*/

insert into SC values('200215121','1',92);
insert into SC values('200215121','2',85);
insert into SC values('200215121','3',88);
insert into SC values('200215122','2',90);
insert into SC values('200215122','3',80);
/*为表 SC 添加数据*/
 
