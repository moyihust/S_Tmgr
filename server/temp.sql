SELECT Student.Sno, Student.Sname, Student.Ssex, Student.Sage, Student.Sdept, Student.Scholarship, Course.Cno, Course.Cname, SC.Grade 
FROM Student
LEFT JOIN SC ON Student.Sno = SC.Sno
LEFT JOIN Course ON SC.Cno = Course.Cno
WHERE Student.Sno = '202112054'; UPDATE Student SET Sage =23 WHERE Sno='202112034';

-- 202112054'; UPDATE Student SET Sage =23 WHERE Sno='202112034
