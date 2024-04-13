CREATE TABLE `S_T`.`Student` (
`Sno
` char(9),
`Sname` char(20) NULL,
`Ssex` char(2) NULL,
`Sage` smallint NULL,
`Sdept
` char(20) NULL,
`Scholarship` char(2) NULL,
CONSTRAINT  PRIMARY KEY (`Sno
`),
CONSTRAINT  UNIQUE (`Sname`)
) DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci ;