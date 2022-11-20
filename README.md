# multicolumn-index-test
mysql多行索引实验。

## 创建实验表

    CREATE TABLE students 
    (id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT, height TINYINT UNSIGNED NOT NULL, gene BIGINT UNSIGNED NOT NULL, name VARCHAR(1024));

    go run main.go

生成一个 500 万行的 students 表，其中 height 有 50 种取值，gene 有 300 万种取值。gene 比 height 更加 selective。
将 students 复制到 students_2。

    CREATE TABLE students_2 AS SELECT * FROM students;
    ALTER TABLE students_2 ADD PRIMARY KEY (id)

分别加上多行索引：

    CREATE INDEX idx_height_gene ON students (height, gene);
    CREATE INDEX idx_gene_height ON students_2 (gene, height);

把 profiling 打开，可以更好地观察每个 query 的耗时。

    SET profiling = 1;

## case 1, (=, =)

    SELECT * FROM students WHERE height = 180 AND gene = 1000;
    SELECT * FROM students_2 WHERE height = 180 AND gene = 1000;
    SHOW profiles;

|211|0.00143700|SELECT * FROM students WHERE height = 180 AND gene = 1000|

|212|0.00064300|SELECT * FROM students_2 WHERE height = 180 AND gene = 1000|

## case 2, (>, >)

213|1.21054800|SELECT * FROM students WHERE height > 180 AND gene > 2400000|
|214|1.18806100|SELECT * FROM students_2 WHERE height > 180 AND gene > 2400000|

## case 3, (=, >)

|215|0.14767000|SELECT * FROM students WHERE height = 180 AND gene > 2400000|
|216|1.05852500|SELECT * FROM students_2 WHERE height = 180 AND gene > 2400000|

## 结论

mysql offical doc: https://dev.mysql.com/doc/refman/8.0/en/range-optimization.html

