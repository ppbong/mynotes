# JPA

## annotation

- `@Entity` 实体
- `@Table` 数据库表
- `@Id` 表主键
- `@GeneratedValue` 主键生成方式
- `@Column` 表字段
- `@Transient` 表外字段（不映射到数据库）
- `@Temporal` 时间类字段
- `@Enumerated` 枚举类字段

### example

```java
public enum Gender {
    MALE, 
    FEMALE
}

@Entity
@Table(name="STUDENT", schema="SCHOOL")
public class Student {
    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    private Long id;
    
    @Column(name="STUDENT_NAME", length=50, nullable=false, unique=false)
    private String name;
    
    @Transient
    private Integer age;
    
    @Temporal(TemporalType.DATE)
    private Date birthDate;
    
    @Enumerated(EnumType.STRING)
    private Gender gender;
    
    // other fields, getters and setters
}
``` 

### TemporalType

```java
@Basic
private java.sql.Date sqlDate;

@Basic
private java.sql.Time sqlTime;

@Basic
private java.sql.Timestamp sqlTimestamp;

@Temporal(TemporalType.DATE)
private java.util.Date utilDate;

@Temporal(TemporalType.TIME)
private java.util.Date utilTime;

@Temporal(TemporalType.TIMESTAMP)
private java.util.Date utilTimestamp;
```

```java
setSqlDate(java.sql.Date.valueOf("2017-11-15"));
setSqlTime(java.sql.Time.valueOf("15:30:14"));
setSqlTimestamp(java.sql.Timestamp.valueOf("2017-11-15 15:30:14.332));

setUtilDate(new SimpleDateFormat("yyyy-MM-dd").parse("2017-11-15"));
setUtilTime(new SimpleDateFormat("HH:mm:ss").parse("15:30:14"));
setUtilTimestamp(new SimpleDateFormat("yyyy-MM-dd HH:mm:ss.SSS").parse("2017-11-15 15:30:14.332"));
```
