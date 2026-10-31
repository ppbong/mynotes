# MyBatis 技术文档

## 1. 介绍

### 1.1 什么是 MyBatis

MyBatis 是一个优秀的持久层框架，它支持自定义 SQL、存储过程以及高级映射。MyBatis 免除了几乎所有的 JDBC 代码和手动设置参数以及获取结果集的工作。MyBatis 可以通过简单的 XML 或注解来配置和映射原始类型、接口和 Java POJO（Plain Old Java Objects，普通老式 Java 对象）为数据库中的记录。

MyBatis 本是 Apache 的一个开源项目 iBatis，2010 年这个项目由 Apache Software Foundation 迁移到了 Google Code，并更名为 MyBatis。2013 年 11 月迁移到 GitHub。

### 1.2 MyBatis 的优势

1. **简单易学**：MyBatis 的架构简单，几乎没有第三方依赖，学习成本较低。
2. **灵活**：MyBatis 不会对应用程序或者数据库的现有设计强加任何影响，SQL 写在 XML 里，便于统一管理和优化。
3. **解除 SQL 与程序代码的耦合**：通过提供 DAO 层，将业务逻辑和数据访问逻辑分离，使系统的设计更清晰，更易维护，更易单元测试。
4. **提供映射标签**：支持对象与数据库的 ORM 字段关系映射。
5. **提供对象关系映射标签**：支持对象关系组件维护。
6. **提供 XML 标签**：支持编写动态 SQL。
7. **支持存储过程**：MyBatis 提供了对存储过程的支持。
8. **与主流框架集成**：MyBatis 可以与 Spring、Spring Boot 等主流框架无缝集成。

### 1.3 MyBatis 与其他持久层框架的对比

#### 1.3.1 MyBatis vs Hibernate

| 特性 | MyBatis | Hibernate |
|------|---------|-----------|
| 定位 | 半 ORM 框架，SQL 映射 | 全 ORM 框架，对象映射 |
| SQL 控制 | 完全控制 SQL，可以优化性能 | 自动生成 SQL，难以优化 |
| 学习曲线 | 简单，容易上手 | 复杂，需要掌握 HQL |
| 灵活性 | 高度灵活，适合复杂查询 | 封装程度高，灵活性较低 |
| 性能 | 性能较好，适合高并发场景 | 性能一般，大量数据时可能存在问题 |
| 适用场景 | 对 SQL 有特殊要求，需要优化性能的项目 | 快速开发，对 SQL 要求不高的项目 |

#### 1.3.2 MyBatis vs JPA

| 特性 | MyBatis | JPA |
|------|---------|-----|
| 规范 | 无标准规范，自定标准 | Java EE 标准规范 |
| SQL 控制 | 直接控制 SQL | 通过 JPQL 或 Criteria API |
| 灵活性 | 更高的灵活性，适合复杂查询 | 标准化程度高，但灵活性相对较低 |
| 学习成本 | 较低 | 中等，需要了解 JPQL |
| 社区支持 | 活跃的社区支持 | 广泛的社区支持，多厂商实现 |

## 2. 安装与配置

### 2.1 环境要求

- JDK 1.8 或更高版本
- Maven 或 Gradle（用于依赖管理）
- 数据库（如 MySQL、Oracle、PostgreSQL 等）

### 2.2 Maven 依赖

在 Maven 项目中，添加 MyBatis 依赖：

```xml
<dependency>
    <groupId>org.mybatis</groupId>
    <artifactId>mybatis</artifactId>
    <version>3.5.15</version>
</dependency>
```

同时，需要添加数据库驱动依赖，以 MySQL 为例：

```xml
<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>8.0.33</version>
</dependency>
```

### 2.3 Gradle 依赖

在 Gradle 项目中，添加 MyBatis 依赖：

```groovy
dependencies {
    implementation 'org.mybatis:mybatis:3.5.15'
    implementation 'mysql:mysql-connector-java:8.0.33'
}
```

### 2.4 基本配置文件

MyBatis 的核心配置文件通常命名为 `mybatis-config.xml`，放在类路径下：

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE configuration
        PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "https://mybatis.org/dtd/mybatis-3-config.dtd">
<configuration>
    <!-- 配置环境 -->
    <environments default="development">
        <environment id="development">
            <!-- 使用 JDBC 事务管理 -->
            <transactionManager type="JDBC"/>
            <!-- 配置数据源 -->
            <dataSource type="POOLED">
                <property name="driver" value="com.mysql.cj.jdbc.Driver"/>
                <property name="url" value="jdbc:mysql://localhost:3306/mybatis_demo?useSSL=false&serverTimezone=UTC"/>
                <property name="username" value="root"/>
                <property name="password" value="password"/>
            </dataSource>
        </environment>
    </environments>
    <!-- 映射文件 -->
    <mappers>
        <mapper resource="mappers/UserMapper.xml"/>
    </mappers>
</configuration>
```

### 2.5 日志配置

MyBatis 可以集成多种日志框架，如 Log4j、Log4j2、Logback 等。以下是集成 Log4j 的配置示例：

首先，添加 Log4j 依赖：

```xml
<dependency>
    <groupId>log4j</groupId>
    <artifactId>log4j</artifactId>
    <version>1.2.17</version>
</dependency>
```

然后，创建 Log4j 配置文件 `log4j.properties`：

```properties
# Global logging configuration
log4j.rootLogger=DEBUG, stdout
# Console output...
log4j.appender.stdout=org.apache.log4j.ConsoleAppender
log4j.appender.stdout.layout=org.apache.log4j.PatternLayout
log4j.appender.stdout.layout.ConversionPattern=%5p [%t] - %m%n

# MyBatis logging configuration
log4j.logger.org.apache.ibatis=DEBUG
log4j.logger.java.sql=DEBUG
```

## 3. 核心概念

### 3.1 SqlSessionFactory

SqlSessionFactory 是 MyBatis 的核心接口，用于创建 SqlSession 实例。SqlSessionFactory 一旦创建，就应该在应用程序的运行期间一直存在，没有任何理由丢弃它或重新创建另一个实例。

通常，SqlSessionFactory 是通过 SqlSessionFactoryBuilder 来创建的：

```java
import org.apache.ibatis.io.Resources;
import org.apache.ibatis.session.SqlSessionFactory;
import org.apache.ibatis.session.SqlSessionFactoryBuilder;

import java.io.IOException;
import java.io.InputStream;

public class MyBatisUtils {
    private static SqlSessionFactory sqlSessionFactory;

    static {
        try {
            // 读取配置文件
            String resource = "mybatis-config.xml";
            InputStream inputStream = Resources.getResourceAsStream(resource);
            // 创建 SqlSessionFactory
            sqlSessionFactory = new SqlSessionFactoryBuilder().build(inputStream);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    // 获取 SqlSession 的静态方法
    public static SqlSessionFactory getSqlSessionFactory() {
        return sqlSessionFactory;
    }
}
```

### 3.2 SqlSession

SqlSession 是 MyBatis 中用于执行持久化操作的对象，它提供了增删改查等方法。SqlSession 实例不是线程安全的，因此每个线程都应该有自己的 SqlSession 实例，并且在使用完毕后关闭它。

```java
import org.apache.ibatis.session.SqlSession;

public class UserDaoImpl implements UserDao {
    private SqlSessionFactory sqlSessionFactory;

    public UserDaoImpl(SqlSessionFactory sqlSessionFactory) {
        this.sqlSessionFactory = sqlSessionFactory;
    }

    @Override
    public User getUserById(int id) {
        SqlSession sqlSession = null;
        try {
            sqlSession = sqlSessionFactory.openSession();
            return sqlSession.selectOne("com.example.mapper.UserMapper.getUserById", id);
        } finally {
            if (sqlSession != null) {
                sqlSession.close();
            }
        }
    }
}
```

### 3.3 Mapper 接口

Mapper 接口是 MyBatis 中用于映射 SQL 语句的接口，它的方法名与 XML 配置文件中的 SQL 语句的 id 相对应。通过 Mapper 接口，我们可以调用对应的 SQL 语句，而不需要编写实现类。

```java
public interface UserMapper {
    User getUserById(int id);
    List<User> getAllUsers();
    int insertUser(User user);
    int updateUser(User user);
    int deleteUser(int id);
}
```

### 3.4 映射文件

映射文件是 MyBatis 中用于配置 SQL 语句和映射关系的 XML 文件，它通常与 Mapper 接口相对应。

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "https://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="com.example.mapper.UserMapper">
    <select id="getUserById" resultType="com.example.pojo.User">
        SELECT * FROM user WHERE id = #{id}
    </select>
    
    <select id="getAllUsers" resultType="com.example.pojo.User">
        SELECT * FROM user
    </select>
    
    <insert id="insertUser" parameterType="com.example.pojo.User">
        INSERT INTO user (username, password, email) VALUES (#{username}, #{password}, #{email})
    </insert>
    
    <update id="updateUser" parameterType="com.example.pojo.User">
        UPDATE user SET username = #{username}, password = #{password}, email = #{email} WHERE id = #{id}
    </update>
    
    <delete id="deleteUser" parameterType="int">
        DELETE FROM user WHERE id = #{id}
    </delete>
</mapper>
```

### 3.5 POJO (Plain Old Java Objects)

POJO 是普通的 Java 对象，它不依赖于任何框架或接口，用于表示数据库中的记录。

```java
public class User {
    private int id;
    private String username;
    private String password;
    private String email;
    
    // getters and setters
    public int getId() {
        return id;
    }
    
    public void setId(int id) {
        this.id = id;
    }
    
    public String getUsername() {
        return username;
    }
    
    public void setUsername(String username) {
        this.username = username;
    }
    
    public String getPassword() {
        return password;
    }
    
    public void setPassword(String password) {
        this.password = password;
    }
    
    public String getEmail() {
        return email;
    }
    
    public void setEmail(String email) {
        this.email = email;
    }
    
    @Override
    public String toString() {
        return "User{" +
                "id=" + id +
                ", username='" + username + '\'' +
                ", password='" + password + '\'' +
                ", email='" + email + '\'' +
                '}';
    }
}
```

## 4. 配置文件详解

### 4.1 核心配置文件结构

MyBatis 的核心配置文件 `mybatis-config.xml` 的结构如下：

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE configuration
        PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "https://mybatis.org/dtd/mybatis-3-config.dtd">
<configuration>
    <!-- 属性配置 -->
    <properties/>
    
    <!-- 设置 -->
    <settings/>
    
    <!-- 类型别名 -->
    <typeAliases/>
    
    <!-- 类型处理器 -->
    <typeHandlers/>
    
    <!-- 对象工厂 -->
    <objectFactory/>
    
    <!-- 插件 -->
    <plugins/>
    
    <!-- 环境配置 -->
    <environments>
        <environment>
            <transactionManager/>
            <dataSource/>
        </environment>
    </environments>
    
    <!-- 数据库厂商标识 -->
    <databaseIdProvider/>
    
    <!-- 映射文件 -->
    <mappers/>
</configuration>
```

### 4.2 属性（properties）

属性可以在外部进行配置，并可以进行动态替换。你既可以在典型的 Java 属性文件中配置这些属性，也可以在 properties 元素的子元素中设置。

```xml
<properties resource="db.properties">
    <property name="username" value="root"/>
    <property name="password" value="password"/>
</properties>
```

然后，在数据源配置中引用这些属性：

```xml
<dataSource type="POOLED">
    <property name="driver" value="${driver}"/>
    <property name="url" value="${url}"/>
    <property name="username" value="${username}"/>
    <property name="password" value="${password}"/>
</dataSource>
```

### 4.3 设置（settings）

设置是 MyBatis 中极为重要的调整设置，它们会改变 MyBatis 的运行时行为。

```xml
<settings>
    <!-- 指定 MyBatis 所用日志的具体实现 -->
    <setting name="logImpl" value="LOG4J"/>
    
    <!-- 全局地开启或关闭配置文件中的所有映射器已经配置的任何缓存 -->
    <setting name="cacheEnabled" value="true"/>
    
    <!-- 延迟加载的全局开关 -->
    <setting name="lazyLoadingEnabled" value="true"/>
    
    <!-- 是否允许单个语句返回多结果集 -->
    <setting name="multipleResultSetsEnabled" value="true"/>
    
    <!-- 是否使用列标签代替列名 -->
    <setting name="useColumnLabel" value="true"/>
    
    <!-- 是否允许 JDBC 支持自动生成主键 -->
    <setting name="useGeneratedKeys" value="false"/>
    
    <!-- 指定 MyBatis 应如何自动映射列到字段或属性 -->
    <setting name="autoMappingBehavior" value="PARTIAL"/>
    
    <!-- 指定发现自动映射目标未知列（或未知属性类型）的行为 -->
    <setting name="autoMappingUnknownColumnBehavior" value="WARNING"/>
    
    <!-- 配置默认的执行器 -->
    <setting name="defaultExecutorType" value="SIMPLE"/>
    
    <!-- 设置超时时间，它决定驱动等待数据库响应的秒数 -->
    <setting name="defaultStatementTimeout" value="25"/>
    
    <!-- 为驱动的结果集获取数量设置一个提示值 -->
    <setting name="defaultFetchSize" value="100"/>
    
    <!-- 是否允许在嵌套语句中使用分页 -->
    <setting name="safeRowBoundsEnabled" value="false"/>
    
    <!-- 是否开启驼峰命名自动映射 -->
    <setting name="mapUnderscoreToCamelCase" value="false"/>
    
    <!-- 是否开启本地缓存 -->
    <setting name="localCacheScope" value="SESSION"/>
    
    <!-- 当没有为参数提供特定的 JDBC 类型时，为空值指定 JDBC 类型 -->
    <setting name="jdbcTypeForNull" value="OTHER"/>
    
    <!-- 指定对象的哪些方法触发一次延迟加载 -->
    <setting name="lazyLoadTriggerMethods" value="equals,clone,hashCode,toString"/>
</settings>
```

### 4.4 类型别名（typeAliases）

类型别名是为 Java 类型设置一个短的名字，可以减少类完全限定名的冗余。

```xml
<typeAliases>
    <!-- 单个类型别名 -->
    <typeAlias type="com.example.pojo.User" alias="User"/>
    
    <!-- 包扫描 -->
    <package name="com.example.pojo"/>
</typeAliases>
```

使用类型别名后，可以在映射文件中简化类名：

```xml
<select id="getUserById" resultType="User">
    SELECT * FROM user WHERE id = #{id}
</select>
```

### 4.5 类型处理器（typeHandlers）

MyBatis 在设置预处理语句（PreparedStatement）中的参数或从结果集中取出一个值时，都会用类型处理器将获取的值以合适的方式转换成 Java 类型。

```xml
<typeHandlers>
    <!-- 自定义类型处理器 -->
    <typeHandler handler="com.example.handler.MyTypeHandler"/>
    
    <!-- 包扫描 -->
    <package name="com.example.handler"/>
</typeHandlers>
```

### 4.6 插件（plugins）

MyBatis 允许你在映射语句执行过程中的某一点进行拦截调用。默认情况下，MyBatis 允许使用插件来拦截的方法调用包括：

- Executor（update, query, flushStatements, commit, rollback, getTransaction, close, isClosed）
- ParameterHandler（getParameterObject, setParameters）
- ResultSetHandler（handleResultSets, handleOutputParameters）
- StatementHandler（prepare, parameterize, batch, update, query）

```xml
<plugins>
    <!-- 分页插件 -->
    <plugin interceptor="com.github.pagehelper.PageInterceptor">
        <!-- 使用下面的方式配置参数，后面会有所有的参数介绍 -->
        <property name="param1" value="value1"/>
    </plugin>
</plugins>
```

### 4.7 环境配置（environments）

MyBatis 可以配置成适应多种环境，这种机制有助于将 SQL 映射应用于多种数据库之中。

```xml
<environments default="development">
    <environment id="development">
        <transactionManager type="JDBC">
            <property name="closeConnection" value="true"/>
        </transactionManager>
        <dataSource type="POOLED">
            <property name="driver" value="${driver}"/>
            <property name="url" value="${url}"/>
            <property name="username" value="${username}"/>
            <property name="password" value="${password}"/>
        </dataSource>
    </environment>
    
    <environment id="test">
        <transactionManager type="JDBC"/>
        <dataSource type="POOLED">
            <property name="driver" value="${test.driver}"/>
            <property name="url" value="${test.url}"/>
            <property name="username" value="${test.username}"/>
            <property name="password" value="${test.password}"/>
        </dataSource>
    </environment>
</environments>
```

### 4.8 映射器（mappers）

映射器是 MyBatis 中最核心的组件之一，它负责将 SQL 语句映射到 Java 方法上。

```xml
<mappers>
    <!-- 使用相对于类路径的资源引用 -->
    <mapper resource="mappers/UserMapper.xml"/>
    
    <!-- 使用完全限定资源定位符（URL） -->
    <mapper url="file:///path/to/UserMapper.xml"/>
    
    <!-- 使用映射器接口实现类的完全限定类名 -->
    <mapper class="com.example.mapper.UserMapper"/>
    
    <!-- 将包内的映射器接口实现全部注册为映射器 -->
    <package name="com.example.mapper"/>
</mappers>
```

## 5. SQL 映射文件和映射器接口

### 5.1 映射文件的结构

SQL 映射文件是 MyBatis 中用于配置 SQL 语句和映射关系的 XML 文件，它的基本结构如下：

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "https://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="com.example.mapper.UserMapper">
    <!-- SQL 语句配置 -->
    <select id="getUserById" resultType="User">
        SELECT * FROM user WHERE id = #{id}
    </select>
    
    <!-- 其他 SQL 语句 -->
</mapper>
```

### 5.2 映射文件的元素

#### 5.2.1 select 元素

select 元素用于配置查询语句，它是 MyBatis 中最常用的元素之一。

```xml
<!-- 简单查询 -->
<select id="getUserById" parameterType="int" resultType="User">
    SELECT * FROM user WHERE id = #{id}
</select>

<!-- 带参数的查询 -->
<select id="getUserByName" parameterType="String" resultType="User">
    SELECT * FROM user WHERE username LIKE CONCAT('%', #{username}, '%')
</select>

<!-- 带多个参数的查询 -->
<select id="getUserByCondition" parameterType="map" resultType="User">
    SELECT * FROM user WHERE username LIKE CONCAT('%', #{username}, '%') AND email LIKE CONCAT('%', #{email}, '%')
</select>

<!-- 带返回结果映射的查询 -->
<select id="getUserWithRole" parameterType="int" resultMap="UserWithRoleMap">
    SELECT u.id, u.username, u.email, r.id as role_id, r.name as role_name
    FROM user u
    LEFT JOIN user_role ur ON u.id = ur.user_id
    LEFT JOIN role r ON ur.role_id = r.id
    WHERE u.id = #{id}
</select>
```

#### 5.2.2 insert 元素

insert 元素用于配置插入语句。

```xml
<!-- 简单插入 -->
<insert id="insertUser" parameterType="User">
    INSERT INTO user (username, password, email) VALUES (#{username}, #{password}, #{email})
</insert>

<!-- 插入并返回主键（MySQL） -->
<insert id="insertUserWithId" parameterType="User" useGeneratedKeys="true" keyProperty="id">
    INSERT INTO user (username, password, email) VALUES (#{username}, #{password}, #{email})
</insert>

<!-- 插入并返回主键（Oracle） -->
<insert id="insertUserOracle" parameterType="User">
    <selectKey keyProperty="id" resultType="int" order="BEFORE">
        SELECT SEQ_USER.NEXTVAL FROM DUAL
    </selectKey>
    INSERT INTO user (id, username, password, email) VALUES (#{id}, #{username}, #{password}, #{email})
</insert>
```

#### 5.2.3 update 元素

update 元素用于配置更新语句。

```xml
<!-- 简单更新 -->
<update id="updateUser" parameterType="User">
    UPDATE user SET username = #{username}, password = #{password}, email = #{email} WHERE id = #{id}
</update>

<!-- 条件更新 -->
<update id="updateUserSelective" parameterType="User">
    UPDATE user
    <set>
        <if test="username != null and username != ''">
            username = #{username},
        </if>
        <if test="password != null and password != ''">
            password = #{password},
        </if>
        <if test="email != null and email != ''">
            email = #{email},
        </if>
    </set>
    WHERE id = #{id}
</update>
```

#### 5.2.4 delete 元素

delete 元素用于配置删除语句。

```xml
<!-- 简单删除 -->
<delete id="deleteUser" parameterType="int">
    DELETE FROM user WHERE id = #{id}
</delete>

<!-- 批量删除 -->
<delete id="deleteUserByIds" parameterType="list">
    DELETE FROM user WHERE id IN
    <foreach collection="list" item="id" open="(" separator="," close=")">
        #{id}
    </foreach>
</delete>
```

#### 5.2.5 resultMap 元素

resultMap 元素用于配置结果集的映射关系，它是 MyBatis 中最强大的元素之一。

```xml
<!-- 基本结果映射 -->
<resultMap id="UserMap" type="User">
    <id property="id" column="id"/>
    <result property="username" column="username"/>
    <result property="password" column="password"/>
    <result property="email" column="email"/>
</resultMap>

<!-- 关联查询结果映射 -->
<resultMap id="UserWithRoleMap" type="User">
    <id property="id" column="id"/>
    <result property="username" column="username"/>
    <result property="email" column="email"/>
    <!-- 一对一关联 -->
    <association property="role" javaType="Role">
        <id property="id" column="role_id"/>
        <result property="name" column="role_name"/>
    </association>
</resultMap>

<!-- 嵌套查询结果映射 -->
<resultMap id="UserWithRoleNestedMap" type="User">
    <id property="id" column="id"/>
    <result property="username" column="username"/>
    <result property="email" column="email"/>
    <!-- 一对一嵌套查询 -->
    <association property="role" javaType="Role" select="getRoleById" column="role_id"/>
</resultMap>

<!-- 一对多关联结果映射 -->
<resultMap id="UserWithOrdersMap" type="User">
    <id property="id" column="id"/>
    <result property="username" column="username"/>
    <result property="email" column="email"/>
    <!-- 一对多关联 -->
    <collection property="orders" ofType="Order">
        <id property="id" column="order_id"/>
        <result property="orderNo" column="order_no"/>
        <result property="amount" column="amount"/>
    </collection>
</resultMap>
```

#### 5.2.6 sql 元素

sql 元素用于定义可重用的 SQL 片段，它可以被其他元素引用。

```xml
<!-- 定义 SQL 片段 -->
<sql id="userColumns">
    id, username, password, email
</sql>

<!-- 引用 SQL 片段 -->
<select id="getUserById" parameterType="int" resultType="User">
    SELECT <include refid="userColumns"/> FROM user WHERE id = #{id}
</select>

<!-- 带参数的 SQL 片段 -->
<sql id="userCondition">
    <if test="username != null and username != ''">
        AND username LIKE CONCAT('%', #{username}, '%')
    </if>
    <if test="email != null and email != ''">
        AND email LIKE CONCAT('%', #{email}, '%')
    </if>
</sql>

<!-- 引用带参数的 SQL 片段 -->
<select id="getUserByCondition" parameterType="map" resultType="User">
    SELECT <include refid="userColumns"/> FROM user WHERE 1=1
    <include refid="userCondition"/>
</select>
```

### 5.3 映射器接口

映射器接口是 MyBatis 中用于映射 SQL 语句的接口，它的方法名与 XML 配置文件中的 SQL 语句的 id 相对应。

#### 5.3.1 基本映射器接口

```java
public interface UserMapper {
    // 根据 ID 查询用户
    User getUserById(int id);
    
    // 查询所有用户
    List<User> getAllUsers();
    
    // 根据用户名模糊查询用户
    List<User> getUserByName(String username);
    
    // 插入用户
    int insertUser(User user);
    
    // 更新用户
    int updateUser(User user);
    
    // 删除用户
    int deleteUser(int id);
}
```

#### 5.3.2 带参数的映射器接口

```java
public interface UserMapper {
    // 使用 @Param 注解指定参数名
    List<User> getUserByCondition(@Param("username") String username, @Param("email") String email);
    
    // 使用 Map 传递多个参数
    List<User> getUserByMap(Map<String, Object> params);
    
    // 使用对象传递多个参数
    List<User> getUserByUserQuery(UserQuery query);
    
    // 批量插入
    int batchInsertUsers(List<User> users);
    
    // 批量更新
    int batchUpdateUsers(List<User> users);
    
    // 批量删除
    int batchDeleteUsers(List<Integer> ids);
}
```

#### 5.3.3 带注解的映射器接口

MyBatis 也支持使用注解来配置 SQL 语句，这样可以不需要 XML 配置文件。

```java
import org.apache.ibatis.annotations.*;

@Mapper
public interface UserMapper {
    // 查询操作
    @Select("SELECT * FROM user WHERE id = #{id}")
    User getUserById(int id);
    
    // 多条件查询
    @Select("SELECT * FROM user WHERE username LIKE CONCAT('%', #{username}, '%') AND email LIKE CONCAT('%', #{email}, '%')")
    List<User> getUserByCondition(@Param("username") String username, @Param("email") String email);
    
    // 插入操作
    @Insert("INSERT INTO user (username, password, email) VALUES (#{username}, #{password}, #{email})")
    @Options(useGeneratedKeys = true, keyProperty = "id")
    int insertUser(User user);
    
    // 更新操作
    @Update("UPDATE user SET username = #{username}, password = #{password}, email = #{email} WHERE id = #{id}")
    int updateUser(User user);
    
    // 删除操作
    @Delete("DELETE FROM user WHERE id = #{id}")
    int deleteUser(int id);
    
    // 批量删除
    @Delete("<script>DELETE FROM user WHERE id IN <foreach collection='list' item='id' open='(' separator=',' close=')'>#{id}</foreach></script>")
    int batchDeleteUsers(List<Integer> ids);
}
```

## 6. 动态 SQL

动态 SQL 是 MyBatis 的强大特性之一，它允许你根据条件动态地构建 SQL 语句，这大大提高了 SQL 语句的灵活性和复用性。

### 6.1 if 元素

if 元素用于条件判断，它可以根据条件动态地包含或排除 SQL 片段。

```xml
<select id="getUserByCondition" parameterType="map" resultType="User">
    SELECT * FROM user WHERE 1=1
    <if test="username != null and username != ''">
        AND username LIKE CONCAT('%', #{username}, '%')
    </if>
    <if test="email != null and email != ''">
        AND email LIKE CONCAT('%', #{email}, '%')
    </if>
    <if test="id != null">
        AND id = #{id}
    </if>
</select>
```

### 6.2 choose、when、otherwise 元素

choose、when、otherwise 元素类似于 Java 中的 switch-case 语句，它们可以根据条件选择其中一个 SQL 片段。

```xml
<select id="getUserByCondition" parameterType="map" resultType="User">
    SELECT * FROM user WHERE 1=1
    <choose>
        <when test="id != null">
            AND id = #{id}
        </when>
        <when test="username != null and username != ''">
            AND username LIKE CONCAT('%', #{username}, '%')
        </when>
        <when test="email != null and email != ''">
            AND email LIKE CONCAT('%', #{email}, '%')
        </when>
        <otherwise>
            AND 1=1
        </otherwise>
    </choose>
</select>
```

### 6.3 trim、where、set 元素

#### 6.3.1 where 元素

where 元素用于处理 SQL 语句中的 WHERE 条件，它可以智能地处理 AND 或 OR 前缀。

```xml
<select id="getUserByCondition" parameterType="map" resultType="User">
    SELECT * FROM user
    <where>
        <if test="username != null and username != ''">
            username LIKE CONCAT('%', #{username}, '%')
        </if>
        <if test="email != null and email != ''">
            AND email LIKE CONCAT('%', #{email}, '%')
        </if>
        <if test="id != null">
            AND id = #{id}
        </if>
    </where>
</select>
```

#### 6.3.2 set 元素

set 元素用于处理 SQL 语句中的 SET 子句，它可以智能地处理逗号。

```xml
<update id="updateUserSelective" parameterType="User">
    UPDATE user
    <set>
        <if test="username != null and username != ''">
            username = #{username},
        </if>
        <if test="password != null and password != ''">
            password = #{password},
        </if>
        <if test="email != null and email != ''">
            email = #{email},
        </if>
    </set>
    WHERE id = #{id}
</update>
```

#### 6.3.3 trim 元素

trim 元素是一个更通用的元素，它可以自定义处理前缀、后缀以及内容。

```xml
<!-- 替代 where 元素 -->
<select id="getUserByCondition" parameterType="map" resultType="User">
    SELECT * FROM user
    <trim prefix="WHERE" prefixOverrides="AND |OR ">
        <if test="username != null and username != ''">
            AND username LIKE CONCAT('%', #{username}, '%')
        </if>
        <if test="email != null and email != ''">
            AND email LIKE CONCAT('%', #{email}, '%')
        </if>
    </trim>
</select>

<!-- 替代 set 元素 -->
<update id="updateUserSelective" parameterType="User">
    UPDATE user
    <trim prefix="SET" suffixOverrides=",">
        <if test="username != null and username != ''">
            username = #{username},
        </if>
        <if test="password != null and password != ''">
            password = #{password},
        </if>
        <if test="email != null and email != ''">
            email = #{email},
        </if>
    </trim>
    WHERE id = #{id}
</update>
```

### 6.4 foreach 元素

foreach 元素用于遍历集合或数组，它可以用于构建 IN 条件、批量插入、批量更新等操作。

```xml
<!-- 构建 IN 条件 -->
<select id="getUserByIds" parameterType="list" resultType="User">
    SELECT * FROM user WHERE id IN
    <foreach collection="list" item="id" index="index" open="(" separator="," close=")">
        #{id}
    </foreach>
</select>

<!-- 批量插入 -->
<insert id="batchInsertUsers" parameterType="list">
    INSERT INTO user (username, password, email) VALUES
    <foreach collection="list" item="user" separator=",">
        (#{user.username}, #{user.password}, #{user.email})
    </foreach>
</insert>

<!-- 批量更新 -->
<update id="batchUpdateUsers" parameterType="list">
    <foreach collection="list" item="user" separator=";">
        UPDATE user SET username = #{user.username}, password = #{user.password}, email = #{user.email} WHERE id = #{user.id}
    </foreach>
</update>

<!-- 批量删除 -->
<delete id="batchDeleteUsers" parameterType="list">
    DELETE FROM user WHERE id IN
    <foreach collection="list" item="id" open="(" separator="," close=")">
        #{id}
    </foreach>
</delete>
```

### 6.5 bind 元素

bind 元素用于在 SQL 语句中创建一个变量，它可以用于简化 SQL 语句中的表达式。

```xml
<select id="getUserByName" parameterType="String" resultType="User">
    <bind name="pattern" value="'%' + username + '%'"/>
    SELECT * FROM user WHERE username LIKE #{pattern}
</select>
```

### 6.6 动态 SQL 的实际应用

#### 6.6.1 高级查询

```xml
<select id="getUserByAdvancedQuery" parameterType="UserQuery" resultType="User">
    SELECT <include refid="userColumns"/> FROM user
    <where>
        <!-- 基本条件 -->
        <if test="username != null and username != ''">
            AND username LIKE CONCAT('%', #{username}, '%')
        </if>
        <if test="email != null and email != ''">
            AND email LIKE CONCAT('%', #{email}, '%')
        </if>
        
        <!-- 范围查询 -->
        <if test="startTime != null">
            AND create_time >= #{startTime}
        </if>
        <if test="endTime != null">
            AND create_time <= #{endTime}
        </if>
        
        <!-- 列表查询 -->
        <if test="ids != null and ids.size() > 0">
            AND id IN
            <foreach collection="ids" item="id" open="(" separator="," close=")">
                #{id}
            </foreach>
        </if>
        
        <!-- 排序 -->
        <if test="orderBy != null and orderBy != ''">
            ORDER BY ${orderBy}
            <if test="orderType != null and orderType != ''">
                ${orderType}
            </if>
        </if>
    </where>
</select>
```

#### 6.6.2 动态更新

```xml
<update id="updateUserSelective" parameterType="User">
    UPDATE user
    <set>
        <if test="username != null and username != ''">
            username = #{username},
        </if>
        <if test="password != null and password != ''">
            password = #{password},
        </if>
        <if test="email != null and email != ''">
            email = #{email},
        </if>
        <if test="status != null">
            status = #{status},
        </if>
        <if test="lastLoginTime != null">
            last_login_time = #{lastLoginTime},
        </if>
        update_time = NOW()
    </set>
    WHERE id = #{id}
</update>
```

#### 6.6.3 动态插入

```xml
<insert id="insertUserSelective" parameterType="User">
    INSERT INTO user
    <trim prefix="(" suffix=")" suffixOverrides=",">
        <if test="username != null and username != ''">
            username,
        </if>
        <if test="password != null and password != ''">
            password,
        </if>
        <if test="email != null and email != ''">
            email,
        </if>
        <if test="status != null">
            status,
        </if>
        create_time,
        update_time
    </trim>
    VALUES
    <trim prefix="(" suffix=")" suffixOverrides=",">
        <if test="username != null and username != ''">
            #{username},
        </if>
        <if test="password != null and password != ''">
            #{password},
        </if>
        <if test="email != null and email != ''">
            #{email},
        </if>
        <if test="status != null">
            #{status},
        </if>
        NOW(),
        NOW()
    </trim>
</insert>
```

## 7. MyBatis 高级特性

### 7.1 缓存机制

MyBatis 提供了强大的缓存机制，可以显著提高查询性能。MyBatis 支持两级缓存：一级缓存（SqlSession 缓存）和二级缓存（namespace 缓存）。

#### 7.1.1 一级缓存

一级缓存是 SqlSession 级别的缓存，它是 MyBatis 默认开启的。当使用同一个 SqlSession 对象执行相同的查询时，MyBatis 会将第一次查询的结果缓存起来，后续的相同查询会直接从缓存中获取结果。

```java
SqlSession sqlSession = sqlSessionFactory.openSession();
UserMapper userMapper = sqlSession.getMapper(UserMapper.class);

// 第一次查询，从数据库获取数据
User user1 = userMapper.getUserById(1);

// 第二次查询，从一级缓存获取数据
User user2 = userMapper.getUserById(1);

// user1 和 user2 是同一个对象
System.out.println(user1 == user2); // true

// 关闭 SqlSession，一级缓存也会被清除
sqlSession.close();
```

#### 7.1.2 二级缓存

二级缓存是 namespace 级别的缓存，它需要手动开启。当使用不同的 SqlSession 对象执行相同的查询时，MyBatis 会将第一次查询的结果缓存起来，后续的相同查询会直接从缓存中获取结果。

##### 开启二级缓存

1. 在 MyBatis 配置文件中开启二级缓存：

```xml
<settings>
    <!-- 开启二级缓存 -->
    <setting name="cacheEnabled" value="true"/>
</settings>
```

2. 在映射文件中配置二级缓存：

```xml
<mapper namespace="com.example.mapper.UserMapper">
    <!-- 配置二级缓存 -->
    <cache 
        eviction="LRU" 
        flushInterval="60000" 
        size="512" 
        readOnly="true"/>
    
    <!-- SQL 语句配置 -->
    <select id="getUserById" parameterType="int" resultType="User" useCache="true">
        SELECT * FROM user WHERE id = #{id}
    </select>
</mapper>
```

##### 缓存属性说明

- `eviction`：缓存回收策略，默认是 LRU（Least Recently Used，最近最少使用）
  - LRU：最近最少使用的，移除最长时间不被使用的对象
  - FIFO：先进先出，按对象进入缓存的顺序来移除它们
  - SOFT：软引用，基于垃圾回收器状态和软引用规则移除对象
  - WEAK：弱引用，更积极地基于垃圾回收器状态和弱引用规则移除对象
- `flushInterval`：缓存刷新间隔，单位是毫秒
- `size`：缓存最大对象数
- `readOnly`：是否只读

##### 使用二级缓存

```java
// 第一个 SqlSession
SqlSession sqlSession1 = sqlSessionFactory.openSession();
UserMapper userMapper1 = sqlSession1.getMapper(UserMapper.class);
User user1 = userMapper1.getUserById(1);
sqlSession1.close(); // 关闭 SqlSession，数据会被保存到二级缓存

// 第二个 SqlSession
SqlSession sqlSession2 = sqlSessionFactory.openSession();
UserMapper userMapper2 = sqlSession2.getMapper(UserMapper.class);
User user2 = userMapper2.getUserById(1); // 从二级缓存获取数据
sqlSession2.close();
```

#### 7.1.3 自定义缓存

MyBatis 允许使用自定义缓存实现，只需要实现 `org.apache.ibatis.cache.Cache` 接口即可。

```java
import org.apache.ibatis.cache.Cache;
import java.util.HashMap;
import java.util.Map;

public class CustomCache implements Cache {
    private final String id;
    private final Map<Object, Object> cache = new HashMap<>();
    
    public CustomCache(String id) {
        this.id = id;
    }
    
    @Override
    public String getId() {
        return id;
    }
    
    @Override
    public void putObject(Object key, Object value) {
        cache.put(key, value);
    }
    
    @Override
    public Object getObject(Object key) {
        return cache.get(key);
    }
    
    @Override
    public Object removeObject(Object key) {
        return cache.remove(key);
    }
    
    @Override
    public void clear() {
        cache.clear();
    }
    
    @Override
    public int getSize() {
        return cache.size();
    }
}
```

在映射文件中使用自定义缓存：

```xml
<mapper namespace="com.example.mapper.UserMapper">
    <!-- 使用自定义缓存 -->
    <cache type="com.example.cache.CustomCache"/>
    
    <!-- SQL 语句配置 -->
</mapper>
```

#### 7.1.4 第三方缓存集成

MyBatis 支持与第三方缓存框架集成，如 Ehcache、Redis、Memcached 等。

##### 集成 Ehcache

1. 添加 Ehcache 依赖：

```xml
<dependency>
    <groupId>org.mybatis.caches</groupId>
    <artifactId>mybatis-ehcache</artifactId>
    <version>1.2.1</version>
</dependency>
<dependency>
    <groupId>net.sf.ehcache</groupId>
    <artifactId>ehcache-core</artifactId>
    <version>2.6.11</version>
</dependency>
```

2. 在映射文件中使用 Ehcache：

```xml
<mapper namespace="com.example.mapper.UserMapper">
    <!-- 使用 Ehcache -->
    <cache type="org.mybatis.caches.ehcache.EhcacheCache"/>
    
    <!-- SQL 语句配置 -->
</mapper>
```

3. 创建 Ehcache 配置文件 `ehcache.xml`：

```xml
<ehcache xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:noNamespaceSchemaLocation="http://ehcache.org/ehcache.xsd"
         updateCheck="false">
    <diskStore path="java.io.tmpdir"/>
    <defaultCache
            maxElementsInMemory="10000"
            eternal="false"
            timeToIdleSeconds="120"
            timeToLiveSeconds="120"
            overflowToDisk="true"
            diskSpoolBufferSizeMB="30"
            maxElementsOnDisk="10000000"
            diskPersistent="false"
            diskExpiryThreadIntervalSeconds="120"
            memoryStoreEvictionPolicy="LRU"
    />
</ehcache>
```

### 7.2 插件机制

MyBatis 提供了强大的插件机制，可以在 SQL 执行的各个阶段进行拦截和增强。MyBatis 支持对以下四种接口进行拦截：

- `Executor`：执行器，负责执行 SQL 语句
- `StatementHandler`：语句处理器，负责处理 SQL 语句的预编译和参数设置
- `ParameterHandler`：参数处理器，负责设置 SQL 语句的参数
- `ResultSetHandler`：结果集处理器，负责处理 SQL 语句的执行结果

#### 7.2.1 自定义插件

创建自定义插件需要实现 `org.apache.ibatis.plugin.Interceptor` 接口，并使用 `@Intercepts` 和 `@Signature` 注解来指定要拦截的接口和方法。

```java
import org.apache.ibatis.executor.Executor;
import org.apache.ibatis.mapping.MappedStatement;
import org.apache.ibatis.plugin.*;

import java.util.Properties;

@Intercepts({
        @Signature(
                type = Executor.class,
                method = "update",
                args = {MappedStatement.class, Object.class}
        ),
        @Signature(
                type = Executor.class,
                method = "query",
                args = {MappedStatement.class, Object.class, RowBounds.class, ResultHandler.class}
        )
})
public class LoggingInterceptor implements Interceptor {
    @Override
    public Object intercept(Invocation invocation) throws Throwable {
        // 获取拦截的方法名
        String methodName = invocation.getMethod().getName();
        
        // 记录开始时间
        long startTime = System.currentTimeMillis();
        
        try {
            // 执行原始方法
            return invocation.proceed();
        } finally {
            // 计算执行时间
            long endTime = System.currentTimeMillis();
            System.out.println("Method " + methodName + " executed in " + (endTime - startTime) + " ms");
        }
    }
    
    @Override
    public Object plugin(Object target) {
        // 使用 Plugin 类的 wrap 方法来包装目标对象
        return Plugin.wrap(target, this);
    }
    
    @Override
    public void setProperties(Properties properties) {
        // 设置插件属性
        String property = properties.getProperty("logLevel");
        if (property != null) {
            System.out.println("Log level: " + property);
        }
    }
}
```

#### 7.2.2 注册插件

在 MyBatis 配置文件中注册插件：

```xml
<plugins>
    <plugin interceptor="com.example.plugin.LoggingInterceptor">
        <property name="logLevel" value="DEBUG"/>
    </plugin>
</plugins>
```

#### 7.2.3 插件的执行顺序

当有多个插件时，MyBatis 会按照配置文件中的顺序来执行插件。插件的执行顺序是：

1. 按照配置文件中的顺序，依次对目标对象进行包装
2. 执行方法时，按照包装的逆序来执行插件的 intercept 方法

### 7.3 类型处理器

MyBatis 的类型处理器用于处理 Java 类型和数据库类型之间的转换。MyBatis 内置了大量的类型处理器，同时也支持自定义类型处理器。

#### 7.3.1 内置类型处理器

MyBatis 内置了以下类型处理器：

- 基本数据类型和包装类型
- String 类型
- Date 类型
- BigDecimal 类型
- Boolean 类型
- 数组类型
- 集合类型
- 枚举类型

#### 7.3.2 自定义类型处理器

创建自定义类型处理器需要实现 `org.apache.ibatis.type.TypeHandler` 接口或继承 `org.apache.ibatis.type.BaseTypeHandler` 类。

```java
import org.apache.ibatis.type.BaseTypeHandler;
import org.apache.ibatis.type.JdbcType;

import java.sql.CallableStatement;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.Date;

public class DateTypeHandler extends BaseTypeHandler<Date> {
    @Override
    public void setNonNullParameter(PreparedStatement ps, int i, Date parameter, JdbcType jdbcType) throws SQLException {
        ps.setLong(i, parameter.getTime());
    }
    
    @Override
    public Date getNullableResult(ResultSet rs, String columnName) throws SQLException {
        long time = rs.getLong(columnName);
        return time > 0 ? new Date(time) : null;
    }
    
    @Override
    public Date getNullableResult(ResultSet rs, int columnIndex) throws SQLException {
        long time = rs.getLong(columnIndex);
        return time > 0 ? new Date(time) : null;
    }
    
    @Override
    public Date getNullableResult(CallableStatement cs, int columnIndex) throws SQLException {
        long time = cs.getLong(columnIndex);
        return time > 0 ? new Date(time) : null;
    }
}
```

#### 7.3.3 注册类型处理器

在 MyBatis 配置文件中注册类型处理器：

```xml
<typeHandlers>
    <typeHandler handler="com.example.typehandler.DateTypeHandler" javaType="java.util.Date" jdbcType="BIGINT"/>
</typeHandlers>
```

或者在映射文件中注册类型处理器：

```xml
<mapper namespace="com.example.mapper.UserMapper">
    <resultMap id="UserMap" type="User">
        <id property="id" column="id"/>
        <result property="username" column="username"/>
        <result property="password" column="password"/>
        <result property="email" column="email"/>
        <result property="createTime" column="create_time" typeHandler="com.example.typehandler.DateTypeHandler"/>
    </resultMap>
</mapper>
```

### 7.4 延迟加载

延迟加载是 MyBatis 的一个重要特性，它可以在需要时才加载关联对象，从而提高查询性能。MyBatis 支持对一对一和一对多关联进行延迟加载。

#### 7.4.1 开启延迟加载

在 MyBatis 配置文件中开启延迟加载：

```xml
<settings>
    <!-- 开启延迟加载 -->
    <setting name="lazyLoadingEnabled" value="true"/>
    <!-- 关闭积极加载，使用按需加载 -->
    <setting name="aggressiveLazyLoading" value="false"/>
</settings>
```

#### 7.4.2 配置延迟加载

在映射文件中配置延迟加载：

```xml
<resultMap id="UserWithRoleMap" type="User">
    <id property="id" column="id"/>
    <result property="username" column="username"/>
    <result property="email" column="email"/>
    <!-- 一对一延迟加载 -->
    <association property="role" javaType="Role" select="getRoleById" column="role_id" fetchType="lazy"/>
</resultMap>

<resultMap id="UserWithOrdersMap" type="User">
    <id property="id" column="id"/>
    <result property="username" column="username"/>
    <result property="email" column="email"/>
    <!-- 一对多延迟加载 -->
    <collection property="orders" ofType="Order" select="getOrdersByUserId" column="id" fetchType="lazy"/>
</resultMap>
```

#### 7.4.3 延迟加载的工作原理

MyBatis 的延迟加载是通过动态代理实现的。当查询一个对象时，MyBatis 会返回一个代理对象。当访问代理对象的关联属性时，代理对象会检查该属性是否已经加载，如果没有加载，则会执行相应的查询来加载该属性。

### 7.5 批处理

MyBatis 提供了批处理功能，可以批量执行 SQL 语句，从而提高插入、更新和删除操作的性能。

#### 7.5.1 使用 SqlSession 进行批处理

```java
SqlSession sqlSession = sqlSessionFactory.openSession(ExecutorType.BATCH);
try {
    UserMapper userMapper = sqlSession.getMapper(UserMapper.class);
    
    // 批量插入
    for (int i = 0; i < 1000; i++) {
        User user = new User();
        user.setUsername("user" + i);
        user.setPassword("password" + i);
        user.setEmail("user" + i + "@example.com");
        userMapper.insertUser(user);
        
        // 每 100 条提交一次
        if (i % 100 == 0) {
            sqlSession.commit();
            sqlSession.clearCache();
        }
    }
    
    // 提交剩余的事务
    sqlSession.commit();
} catch (Exception e) {
    sqlSession.rollback();
    e.printStackTrace();
} finally {
    sqlSession.close();
}
```

#### 7.5.2 使用 foreach 进行批量操作

除了使用 SqlSession 进行批处理外，还可以使用 foreach 元素进行批量操作。

```xml
<!-- 批量插入 -->
<insert id="batchInsertUsers" parameterType="list">
    INSERT INTO user (username, password, email) VALUES
    <foreach collection="list" item="user" separator=",">
        (#{user.username}, #{user.password}, #{user.email})
    </foreach>
</insert>

<!-- 批量更新 -->
<update id="batchUpdateUsers" parameterType="list">
    <foreach collection="list" item="user" separator=";">
        UPDATE user SET username = #{user.username}, password = #{user.password}, email = #{user.email} WHERE id = #{user.id}
    </foreach>
</update>

<!-- 批量删除 -->
<delete id="batchDeleteUsers" parameterType="list">
    DELETE FROM user WHERE id IN
    <foreach collection="list" item="id" open="(" separator="," close=")">
        #{id}
    </foreach>
</delete>
```

### 7.6 多数据库支持

MyBatis 支持多种数据库，可以通过数据库厂商标识（databaseId）来为不同的数据库提供不同的 SQL 语句。

#### 7.6.1 配置多数据库支持

在 MyBatis 配置文件中配置多数据库支持：

```xml
<databaseIdProvider type="DB_VENDOR">
    <property name="MySQL" value="mysql"/>
    <property name="Oracle" value="oracle"/>
    <property name="SQL Server" value="sqlserver"/>
</databaseIdProvider>
```

#### 7.6.2 使用数据库厂商标识

在映射文件中使用数据库厂商标识：

```xml
<!-- MySQL 特定的 SQL -->
<select id="getUserById" parameterType="int" resultType="User" databaseId="mysql">
    SELECT * FROM user WHERE id = #{id}
</select>

<!-- Oracle 特定的 SQL -->
<select id="getUserById" parameterType="int" resultType="User" databaseId="oracle">
    SELECT * FROM user WHERE id = #{id}
</select>

<!-- 默认 SQL，当没有匹配的数据库厂商标识时使用 -->
<select id="getUserById" parameterType="int" resultType="User">
    SELECT * FROM user WHERE id = #{id}
</select>
```

### 7.7 存储过程支持

MyBatis 支持调用存储过程，可以使用 `callable` 属性来指定 SQL 语句是存储过程。

#### 7.7.1 调用存储过程

```xml
<!-- 调用存储过程 -->
<select id="getUserByProcedure" parameterType="map" statementType="CALLABLE">
    {CALL getUserById(#{id, mode=IN, jdbcType=INTEGER}, #{user, mode=OUT, jdbcType=CURSOR, resultMap=UserMap})}
</select>
```

在 Java 代码中调用存储过程：

```java
Map<String, Object> params = new HashMap<>();
params.put("id", 1);
userMapper.getUserByProcedure(params);
List<User> users = (List<User>) params.get("user");
```

### 7.8 动态脚本

MyBatis 支持使用动态脚本来构建 SQL 语句，可以使用 `script` 标签来在注解中使用动态 SQL。

```java
import org.apache.ibatis.annotations.*;

@Mapper
public interface UserMapper {
    @Select("<script>SELECT * FROM user WHERE 1=1 <if test='username != null'>AND username LIKE CONCAT('%', #{username}, '%')</if></script>")
    List<User> getUserByCondition(@Param("username") String username);
}
```

## 8. MyBatis 与 Spring 框架集成

MyBatis 与 Spring 框架的集成是企业级应用开发中非常常见的场景。MyBatis 提供了专门的 `mybatis-spring` 模块，用于简化与 Spring 的集成。

### 8.1 基本集成方式

#### 8.1.1 添加依赖

在 Maven 项目中添加以下依赖：

```xml
<!-- MyBatis -->
<dependency>
    <groupId>org.mybatis</groupId>
    <artifactId>mybatis</artifactId>
    <version>3.5.13</version>
</dependency>

<!-- MyBatis Spring 集成 -->
<dependency>
    <groupId>org.mybatis</groupId>
    <artifactId>mybatis-spring</artifactId>
    <version>3.0.3</version>
</dependency>

<!-- Spring JDBC -->
<dependency>
    <groupId>org.springframework</groupId>
    <artifactId>spring-jdbc</artifactId>
    <version>6.0.13</version>
</dependency>

<!-- Spring Context -->
<dependency>
    <groupId>org.springframework</groupId>
    <artifactId>spring-context</artifactId>
    <version>6.0.13</version>
</dependency>

<!-- 数据库驱动 -->
<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>8.0.33</version>
</dependency>

<!-- 数据源 -->
<dependency>
    <groupId>com.zaxxer</groupId>
    <artifactId>HikariCP</artifactId>
    <version>5.0.1</version>
</dependency>
```

#### 8.1.2 配置 Spring

创建 Spring 配置文件 `applicationContext.xml`：

```xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xmlns:tx="http://www.springframework.org/schema/tx"
       xsi:schemaLocation="http://www.springframework.org/schema/beans
       http://www.springframework.org/schema/beans/spring-beans.xsd
       http://www.springframework.org/schema/tx
       http://www.springframework.org/schema/tx/spring-tx.xsd">

    <!-- 数据源配置 -->
    <bean id="dataSource" class="com.zaxxer.hikari.HikariDataSource">
        <property name="driverClassName" value="com.mysql.cj.jdbc.Driver"/>
        <property name="jdbcUrl" value="jdbc:mysql://localhost:3306/mybatis?useSSL=false&serverTimezone=UTC"/>
        <property name="username" value="root"/>
        <property name="password" value="password"/>
        <property name="maximumPoolSize" value="10"/>
    </bean>

    <!-- SqlSessionFactory 配置 -->
    <bean id="sqlSessionFactory" class="org.mybatis.spring.SqlSessionFactoryBean">
        <property name="dataSource" ref="dataSource"/>
        <property name="configLocation" value="classpath:mybatis-config.xml"/>
        <property name="mapperLocations" value="classpath:mappers/*.xml"/>
        <property name="typeAliasesPackage" value="com.example.entity"/>
    </bean>

    <!-- SqlSessionTemplate 配置 -->
    <bean id="sqlSession" class="org.mybatis.spring.SqlSessionTemplate">
        <constructor-arg index="0" ref="sqlSessionFactory"/>
    </bean>

    <!-- Mapper 扫描配置 -->
    <bean class="org.mybatis.spring.mapper.MapperScannerConfigurer">
        <property name="basePackage" value="com.example.mapper"/>
        <property name="sqlSessionFactoryBeanName" value="sqlSessionFactory"/>
    </bean>

    <!-- 事务管理器配置 -->
    <bean id="transactionManager" class="org.springframework.jdbc.datasource.DataSourceTransactionManager">
        <property name="dataSource" ref="dataSource"/>
    </bean>

    <!-- 事务注解支持 -->
    <tx:annotation-driven transaction-manager="transactionManager"/>
</beans>
```

#### 8.1.3 配置 MyBatis

创建 MyBatis 配置文件 `mybatis-config.xml`：

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE configuration
        PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "https://mybatis.org/dtd/mybatis-3-config.dtd">
<configuration>
    <settings>
        <!-- 开启二级缓存 -->
        <setting name="cacheEnabled" value="true"/>
        <!-- 开启延迟加载 -->
        <setting name="lazyLoadingEnabled" value="true"/>
        <!-- 关闭积极加载 -->
        <setting name="aggressiveLazyLoading" value="false"/>
        <!-- 开启驼峰命名转换 -->
        <setting name="mapUnderscoreToCamelCase" value="true"/>
        <!-- 日志配置 -->
        <setting name="logImpl" value="STDOUT_LOGGING"/>
    </settings>
</configuration>
```

#### 8.1.4 创建实体类

```java
public class User {
    private Integer id;
    private String username;
    private String password;
    private String email;
    private Date createTime;
    private Date updateTime;
    
    // getter 和 setter 方法
    // ...
}
```

#### 8.1.5 创建映射器接口

```java
public interface UserMapper {
    User getUserById(Integer id);
    List<User> getAllUsers();
    int insertUser(User user);
    int updateUser(User user);
    int deleteUser(Integer id);
}
```

#### 8.1.6 创建映射文件

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "https://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="com.example.mapper.UserMapper">
    <select id="getUserById" parameterType="int" resultType="User">
        SELECT * FROM user WHERE id = #{id}
    </select>
    
    <select id="getAllUsers" resultType="User">
        SELECT * FROM user
    </select>
    
    <insert id="insertUser" parameterType="User" useGeneratedKeys="true" keyProperty="id">
        INSERT INTO user (username, password, email, create_time, update_time) 
        VALUES (#{username}, #{password}, #{email}, NOW(), NOW())
    </insert>
    
    <update id="updateUser" parameterType="User">
        UPDATE user SET username = #{username}, password = #{password}, email = #{email}, update_time = NOW()
        WHERE id = #{id}
    </update>
    
    <delete id="deleteUser" parameterType="int">
        DELETE FROM user WHERE id = #{id}
    </delete>
</mapper>
```

#### 8.1.7 创建服务层

```java
@Service
public class UserService {
    @Autowired
    private UserMapper userMapper;
    
    @Transactional
    public User getUserById(Integer id) {
        return userMapper.getUserById(id);
    }
    
    @Transactional
    public List<User> getAllUsers() {
        return userMapper.getAllUsers();
    }
    
    @Transactional
    public int insertUser(User user) {
        return userMapper.insertUser(user);
    }
    
    @Transactional
    public int updateUser(User user) {
        return userMapper.updateUser(user);
    }
    
    @Transactional
    public int deleteUser(Integer id) {
        return userMapper.deleteUser(id);
    }
}
```

### 8.2 Spring Boot 集成

Spring Boot 提供了更简单的方式来集成 MyBatis，通过 `mybatis-spring-boot-starter` 依赖可以快速搭建 MyBatis 应用。

#### 8.2.1 添加依赖

在 Maven 项目中添加以下依赖：

```xml
<!-- Spring Boot Starter Web -->
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-web</artifactId>
    <version>3.1.4</version>
</dependency>

<!-- MyBatis Spring Boot Starter -->
<dependency>
    <groupId>org.mybatis.spring.boot</groupId>
    <artifactId>mybatis-spring-boot-starter</artifactId>
    <version>3.0.3</version>
</dependency>

<!-- MySQL 驱动 -->
<dependency>
    <groupId>com.mysql</groupId>
    <artifactId>mysql-connector-j</artifactId>
    <version>8.1.0</version>
    <scope>runtime</scope>
</dependency>

<!-- HikariCP 数据源 -->
<dependency>
    <groupId>com.zaxxer</groupId>
    <artifactId>HikariCP</artifactId>
    <version>5.0.1</version>
</dependency>
```

#### 8.2.2 配置 application.properties

```properties
# 数据库连接配置
spring.datasource.url=jdbc:mysql://localhost:3306/mybatis?useSSL=false&serverTimezone=UTC
spring.datasource.username=root
spring.datasource.password=password
spring.datasource.driver-class-name=com.mysql.cj.jdbc.Driver

# HikariCP 配置
spring.datasource.hikari.maximum-pool-size=10
spring.datasource.hikari.minimum-idle=5
spring.datasource.hikari.idle-timeout=30000
spring.datasource.hikari.max-lifetime=1800000

# MyBatis 配置
mybatis.mapper-locations=classpath:mappers/*.xml
mybatis.type-aliases-package=com.example.entity
mybatis.configuration.map-underscore-to-camel-case=true
mybatis.configuration.cache-enabled=true
mybatis.configuration.lazy-loading-enabled=true
mybatis.configuration.aggressive-lazy-loading=false
mybatis.configuration.log-impl=org.apache.ibatis.logging.stdout.StdOutImpl
```

#### 8.2.3 创建 Spring Boot 主类

```java
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MapperScan("com.example.mapper") // 扫描映射器接口
public class MyBatisSpringBootApplication {
    public static void main(String[] args) {
        SpringApplication.run(MyBatisSpringBootApplication.class, args);
    }
}
```

#### 8.2.4 创建实体类

```java
public class User {
    private Integer id;
    private String username;
    private String password;
    private String email;
    private Date createTime;
    private Date updateTime;
    
    // getter 和 setter 方法
    // ...
}
```

#### 8.2.5 创建映射器接口

```java
@Mapper
public interface UserMapper {
    User getUserById(Integer id);
    List<User> getAllUsers();
    int insertUser(User user);
    int updateUser(User user);
    int deleteUser(Integer id);
}
```

#### 8.2.6 创建映射文件

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "https://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="com.example.mapper.UserMapper">
    <select id="getUserById" parameterType="int" resultType="User">
        SELECT * FROM user WHERE id = #{id}
    </select>
    
    <select id="getAllUsers" resultType="User">
        SELECT * FROM user
    </select>
    
    <insert id="insertUser" parameterType="User" useGeneratedKeys="true" keyProperty="id">
        INSERT INTO user (username, password, email, create_time, update_time) 
        VALUES (#{username}, #{password}, #{email}, NOW(), NOW())
    </insert>
    
    <update id="updateUser" parameterType="User">
        UPDATE user SET username = #{username}, password = #{password}, email = #{email}, update_time = NOW()
        WHERE id = #{id}
    </update>
    
    <delete id="deleteUser" parameterType="int">
        DELETE FROM user WHERE id = #{id}
    </delete>
</mapper>
```

#### 8.2.7 创建服务层

```java
@Service
public class UserService {
    @Autowired
    private UserMapper userMapper;
    
    @Transactional
    public User getUserById(Integer id) {
        return userMapper.getUserById(id);
    }
    
    @Transactional
    public List<User> getAllUsers() {
        return userMapper.getAllUsers();
    }
    
    @Transactional
    public int insertUser(User user) {
        return userMapper.insertUser(user);
    }
    
    @Transactional
    public int updateUser(User user) {
        return userMapper.updateUser(user);
    }
    
    @Transactional
    public int deleteUser(Integer id) {
        return userMapper.deleteUser(id);
    }
}
```

#### 8.2.8 创建控制器

```java
@RestController
@RequestMapping("/api/users")
public class UserController {
    @Autowired
    private UserService userService;
    
    @GetMapping("/{id}")
    public ResponseEntity<User> getUserById(@PathVariable Integer id) {
        User user = userService.getUserById(id);
        if (user == null) {
            return ResponseEntity.notFound().build();
        }
        return ResponseEntity.ok(user);
    }
    
    @GetMapping
    public ResponseEntity<List<User>> getAllUsers() {
        List<User> users = userService.getAllUsers();
        return ResponseEntity.ok(users);
    }
    
    @PostMapping
    public ResponseEntity<User> insertUser(@RequestBody User user) {
        userService.insertUser(user);
        return ResponseEntity.ok(user);
    }
    
    @PutMapping("/{id}")
    public ResponseEntity<User> updateUser(@PathVariable Integer id, @RequestBody User user) {
        user.setId(id);
        userService.updateUser(user);
        return ResponseEntity.ok(user);
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteUser(@PathVariable Integer id) {
        userService.deleteUser(id);
        return ResponseEntity.noContent().build();
    }
}
```

### 8.3 事务管理

MyBatis 与 Spring 集成后，可以使用 Spring 的事务管理功能来管理 MyBatis 的事务。

#### 8.3.1 声明式事务

Spring 提供了声明式事务管理，通过 `@Transactional` 注解可以轻松地为方法添加事务支持。

```java
@Service
public class UserService {
    @Autowired
    private UserMapper userMapper;
    
    @Autowired
    private OrderMapper orderMapper;
    
    // 支持事务的方法
    @Transactional(rollbackFor = Exception.class)
    public void createUserWithOrder(User user, Order order) {
        // 插入用户
        userMapper.insertUser(user);
        
        // 设置订单关联的用户ID
        order.setUserId(user.getId());
        
        // 插入订单
        orderMapper.insertOrder(order);
    }
}
```

#### 8.3.2 事务传播行为

Spring 支持多种事务传播行为，可以通过 `@Transactional` 注解的 `propagation` 属性来设置。

```java
@Service
public class UserService {
    @Autowired
    private UserMapper userMapper;
    
    @Autowired
    private EmailService emailService;
    
    @Transactional(propagation = Propagation.REQUIRED)
    public void createUser(User user) {
        userMapper.insertUser(user);
        
        // 调用其他服务的方法
        emailService.sendWelcomeEmail(user);
    }
}

@Service
public class EmailService {
    @Transactional(propagation = Propagation.REQUIRES_NEW)
    public void sendWelcomeEmail(User user) {
        // 发送邮件的逻辑
        // ...
    }
}
```

#### 8.3.3 事务隔离级别

Spring 支持多种事务隔离级别，可以通过 `@Transactional` 注解的 `isolation` 属性来设置。

```java
@Service
public class UserService {
    @Autowired
    private UserMapper userMapper;
    
    @Transactional(isolation = Isolation.READ_COMMITTED)
    public User getUserById(Integer id) {
        return userMapper.getUserById(id);
    }
}
```

### 8.4 高级配置

#### 8.4.1 自定义 SqlSessionFactory

在 Spring Boot 中，可以通过配置类来自定义 SqlSessionFactory：

```java
@Configuration
public class MyBatisConfig {
    @Bean
    public SqlSessionFactory sqlSessionFactory(DataSource dataSource) throws Exception {
        SqlSessionFactoryBean factoryBean = new SqlSessionFactoryBean();
        factoryBean.setDataSource(dataSource);
        factoryBean.setMapperLocations(new PathMatchingResourcePatternResolver().getResources("classpath:mappers/*.xml"));
        factoryBean.setTypeAliasesPackage("com.example.entity");
        
        // 自定义配置
        org.apache.ibatis.session.Configuration configuration = new org.apache.ibatis.session.Configuration();
        configuration.setMapUnderscoreToCamelCase(true);
        configuration.setCacheEnabled(true);
        configuration.setLazyLoadingEnabled(true);
        configuration.setAggressiveLazyLoading(false);
        factoryBean.setConfiguration(configuration);
        
        return factoryBean.getObject();
    }
}
```

#### 8.4.2 自定义 TypeHandler

在 Spring Boot 中，可以通过配置类来注册自定义的 TypeHandler：

```java
@Configuration
public class MyBatisConfig {
    @Bean
    public SqlSessionFactory sqlSessionFactory(DataSource dataSource) throws Exception {
        SqlSessionFactoryBean factoryBean = new SqlSessionFactoryBean();
        factoryBean.setDataSource(dataSource);
        
        // 注册自定义 TypeHandler
        TypeHandlerRegistry typeHandlerRegistry = factoryBean.getObject().getConfiguration().getTypeHandlerRegistry();
        typeHandlerRegistry.register(Date.class, DateTypeHandler.class);
        
        return factoryBean.getObject();
    }
}
```

#### 8.4.3 自定义插件

在 Spring Boot 中，可以通过配置类来注册自定义的插件：

```java
@Configuration
public class MyBatisConfig {
    @Bean
    public SqlSessionFactory sqlSessionFactory(DataSource dataSource) throws Exception {
        SqlSessionFactoryBean factoryBean = new SqlSessionFactoryBean();
        factoryBean.setDataSource(dataSource);
        
        // 注册自定义插件
        factoryBean.setPlugins(new LoggingInterceptor());
        
        return factoryBean.getObject();
    }
}
```

### 8.5 MyBatis 与 Spring Boot 的高级特性

#### 8.5.1 分页插件集成

MyBatis 可以与分页插件（如 PageHelper）集成，实现分页功能。

1. 添加依赖：

```xml
<dependency>
    <groupId>com.github.pagehelper</groupId>
    <artifactId>pagehelper-spring-boot-starter</artifactId>
    <version>1.4.6</version>
</dependency>
```

2. 配置 application.properties：

```properties
# PageHelper 配置
pagehelper.helper-dialect=mysql
pagehelper.reasonable=true
pagehelper.support-methods-arguments=true
pagehelper.params=count=countSql
```

3. 使用分页：

```java
@Service
public class UserService {
    @Autowired
    private UserMapper userMapper;
    
    public PageInfo<User> getUsersByPage(int pageNum, int pageSize) {
        // 开始分页
        PageHelper.startPage(pageNum, pageSize);
        
        // 查询数据
        List<User> users = userMapper.getAllUsers();
        
        // 包装分页结果
        return new PageInfo<>(users);
    }
}
```

#### 8.5.2 多数据源配置

在 Spring Boot 中，可以配置多个数据源来支持多数据库操作。

1. 添加依赖：

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-jdbc</artifactId>
</dependency>
```

2. 配置 application.properties：

```properties
# 主数据源配置
spring.datasource.primary.url=jdbc:mysql://localhost:3306/mybatis?useSSL=false&serverTimezone=UTC
spring.datasource.primary.username=root
spring.datasource.primary.password=password
spring.datasource.primary.driver-class-name=com.mysql.cj.jdbc.Driver

# 从数据源配置
spring.datasource.secondary.url=jdbc:mysql://localhost:3306/mybatis_secondary?useSSL=false&serverTimezone=UTC
spring.datasource.secondary.username=root
spring.datasource.secondary.password=password
spring.datasource.secondary.driver-class-name=com.mysql.cj.jdbc.Driver
```

3. 配置多数据源：

```java
@Configuration
public class DataSourceConfig {
    // 主数据源配置
    @Bean
    @Primary
    @ConfigurationProperties(prefix = "spring.datasource.primary")
    public DataSource primaryDataSource() {
        return DataSourceBuilder.create().build();
    }
    
    // 从数据源配置
    @Bean
    @ConfigurationProperties(prefix = "spring.datasource.secondary")
    public DataSource secondaryDataSource() {
        return DataSourceBuilder.create().build();
    }
    
    // 主数据源 SqlSessionFactory
    @Bean
    @Primary
    public SqlSessionFactory primarySqlSessionFactory(@Qualifier("primaryDataSource") DataSource dataSource) throws Exception {
        SqlSessionFactoryBean factoryBean = new SqlSessionFactoryBean();
        factoryBean.setDataSource(dataSource);
        factoryBean.setMapperLocations(new PathMatchingResourcePatternResolver().getResources("classpath:mappers/primary/*.xml"));
        factoryBean.setTypeAliasesPackage("com.example.entity");
        return factoryBean.getObject();
    }
    
    // 从数据源 SqlSessionFactory
    @Bean
    public SqlSessionFactory secondarySqlSessionFactory(@Qualifier("secondaryDataSource") DataSource dataSource) throws Exception {
        SqlSessionFactoryBean factoryBean = new SqlSessionFactoryBean();
        factoryBean.setDataSource(dataSource);
        factoryBean.setMapperLocations(new PathMatchingResourcePatternResolver().getResources("classpath:mappers/secondary/*.xml"));
        factoryBean.setTypeAliasesPackage("com.example.entity");
        return factoryBean.getObject();
    }
    
    // 主数据源 SqlSessionTemplate
    @Bean
    @Primary
    public SqlSessionTemplate primarySqlSessionTemplate(@Qualifier("primarySqlSessionFactory") SqlSessionFactory sqlSessionFactory) {
        return new SqlSessionTemplate(sqlSessionFactory);
    }
    
    // 从数据源 SqlSessionTemplate
    @Bean
    public SqlSessionTemplate secondarySqlSessionTemplate(@Qualifier("secondarySqlSessionFactory") SqlSessionFactory sqlSessionFactory) {
        return new SqlSessionTemplate(sqlSessionFactory);
    }
}
```

4. 配置映射器扫描：

```java
@Configuration
@MapperScan(basePackages = "com.example.mapper.primary", sqlSessionTemplateRef = "primarySqlSessionTemplate")
public class PrimaryMapperConfig {
}

@Configuration
@MapperScan(basePackages = "com.example.mapper.secondary", sqlSessionTemplateRef = "secondarySqlSessionTemplate")
public class SecondaryMapperConfig {
}
```

### 8.6 MyBatis-Plus 集成

MyBatis-Plus 是 MyBatis 的增强工具，提供了许多实用的功能，如自动生成 SQL、分页、代码生成等。

#### 8.6.1 添加依赖

```xml
<!-- MyBatis-Plus -->
<dependency>
    <groupId>com.baomidou</groupId>
    <artifactId>mybatis-plus-boot-starter</artifactId>
    <version>3.5.3.1</version>
</dependency>
```

#### 8.6.2 配置 application.properties

```properties
# MyBatis-Plus 配置
mybatis-plus.mapper-locations=classpath:mappers/*.xml
mybatis-plus.type-aliases-package=com.example.entity
mybatis-plus.configuration.map-underscore-to-camel-case=true
mybatis-plus.global-config.db-config.id-type=auto
mybatis-plus.global-config.db-config.logic-delete-field=deleted
mybatis-plus.global-config.db-config.logic-delete-value=1
mybatis-plus.global-config.db-config.logic-not-delete-value=0
```

#### 8.6.3 创建实体类

```java
@Data
@TableName("user")
public class User {
    @TableId(type = IdType.AUTO)
    private Integer id;
    private String username;
    private String password;
    private String email;
    private Date createTime;
    private Date updateTime;
    @TableLogic
    private Integer deleted;
}
```

#### 8.6.4 创建映射器接口

```java
@Mapper
public interface UserMapper extends BaseMapper<User> {
    // 自定义方法
    List<User> getUserByCondition(@Param("username") String username, @Param("email") String email);
}
```

#### 8.6.5 创建服务层

```java
@Service
public class UserService extends ServiceImpl<UserMapper, User> {
    public List<User> getUserByCondition(String username, String email) {
        return baseMapper.getUserByCondition(username, email);
    }
    
    public Page<User> getUserByPage(int pageNum, int pageSize) {
        Page<User> page = new Page<>(pageNum, pageSize);
        return baseMapper.selectPage(page, null);
    }
}
```

#### 8.6.6 使用 MyBatis-Plus

```java
@RestController
@RequestMapping("/api/users")
public class UserController {
    @Autowired
    private UserService userService;
    
    @GetMapping
    public ResponseEntity<Page<User>> getUsersByPage(
            @RequestParam(defaultValue = "1") int pageNum,
            @RequestParam(defaultValue = "10") int pageSize) {
        Page<User> page = userService.getUserByPage(pageNum, pageSize);
        return ResponseEntity.ok(page);
    }
    
    @GetMapping("/condition")
    public ResponseEntity<List<User>> getUserByCondition(
            @RequestParam(required = false) String username,
            @RequestParam(required = false) String email) {
        List<User> users = userService.getUserByCondition(username, email);
        return ResponseEntity.ok(users);
    }
}
```

## 9. MyBatis 最佳实践和性能优化

### 9.1 最佳实践

#### 9.1.1 命名规范

- **映射文件命名**：使用 `Mapper.xml` 作为后缀，如 `UserMapper.xml`
- **映射器接口命名**：使用 `Mapper` 作为后缀，如 `UserMapper.java`
- **SQL 语句 ID 命名**：使用动词+名词的形式，如 `getUserById`、`insertUser`
- **参数命名**：使用有意义的参数名，避免使用 `param1`、`param2` 等

#### 9.1.2 代码组织

- **按功能模块组织代码**：将相关的实体类、映射器、服务层和控制器放在同一个包下
- **分离关注点**：将 SQL 语句与 Java 代码分离，便于维护
- **使用接口**：使用接口定义映射器，便于测试和扩展

#### 9.1.3 SQL 编写

- **使用参数化查询**：避免 SQL 注入攻击
- **使用别名**：使用别名来简化 SQL 语句和映射配置
- **避免 SELECT ***：只查询需要的字段，减少数据传输量
- **使用联表查询**：合理使用联表查询，减少数据库访问次数

#### 9.1.4 事务管理

- **合理设置事务边界**：只在需要事务的方法上添加 `@Transactional` 注解
- **设置合适的事务隔离级别**：根据业务需求选择合适的事务隔离级别
- **设置合适的事务传播行为**：根据业务需求选择合适的事务传播行为

#### 9.1.5 错误处理

- **统一异常处理**：使用全局异常处理器来处理异常
- **记录错误日志**：记录详细的错误日志，便于调试和排查问题
- **提供友好的错误信息**：向用户提供友好的错误信息，避免暴露系统内部错误

### 9.2 性能优化

#### 9.2.1 缓存优化

- **合理使用缓存**：根据数据的更新频率和查询频率来决定是否使用缓存
- **设置合适的缓存过期时间**：避免缓存数据过期导致的数据不一致问题
- **使用第三方缓存**：对于高并发场景，可以使用 Redis 等第三方缓存来提高性能

#### 9.2.2 SQL 优化

- **使用索引**：为经常查询的字段创建索引
- **避免全表扫描**：使用 WHERE 条件来过滤数据，避免全表扫描
- **避免嵌套查询**：尽量使用联表查询替代嵌套查询
- **使用 LIMIT**：限制查询结果的数量，避免返回过多的数据

#### 9.2.3 连接池优化

- **选择合适的连接池**：如 HikariCP、Druid 等
- **设置合适的连接池参数**：根据业务需求设置合适的连接池大小、最大连接数、最小连接数等
- **监控连接池**：监控连接池的使用情况，及时发现和解决问题

#### 9.2.4 批量操作

- **使用批量插入**：对于大量数据的插入操作，使用批量插入可以提高性能
- **使用批量更新**：对于大量数据的更新操作，使用批量更新可以提高性能
- **使用批量删除**：对于大量数据的删除操作，使用批量删除可以提高性能

#### 9.2.5 延迟加载

- **合理使用延迟加载**：对于关联数据，使用延迟加载可以提高查询性能
- **避免 N+1 查询问题**：合理使用关联查询或分批查询来避免 N+1 查询问题

#### 9.2.6 分页优化

- **使用物理分页**：使用数据库的分页功能，避免查询所有数据后再分页
- **设置合理的分页大小**：根据业务需求设置合适的分页大小，避免返回过多的数据
- **使用索引排序**：对于需要排序的分页查询，使用索引可以提高性能

### 9.3 安全考虑

#### 9.3.1 SQL 注入防护

- **使用参数化查询**：避免直接拼接 SQL 语句
- **使用 MyBatis 的 #{} 语法**：使用 #{} 语法来设置参数，MyBatis 会自动进行参数化
- **避免使用 ${} 语法**：如果必须使用 ${} 语法，需要进行严格的参数验证

#### 9.3.2 敏感数据保护

- **加密敏感数据**：对于敏感数据（如密码），使用加密算法进行加密存储
- **脱敏处理**：对于需要展示的敏感数据，进行脱敏处理
- **权限控制**：设置严格的权限控制，确保只有授权用户才能访问敏感数据

#### 9.3.3 防 XSS 攻击

- **输入验证**：对用户输入进行严格的验证
- **输出转义**：对输出的数据进行转义，避免 XSS 攻击
- **使用安全框架**：使用安全框架（如 Spring Security）来防护 XSS 攻击

## 10. MyBatis 常见问题和解决方案

### 10.1 配置问题

#### 10.1.1 找不到映射文件

**问题**：启动应用时出现 `Invalid bound statement (not found)` 错误，提示找不到映射文件。

**解决方案**：
- 检查映射文件的路径是否正确
- 检查映射文件的命名空间是否与映射器接口的全限定名一致
- 检查 SQL 语句的 ID 是否与映射器接口的方法名一致
- 检查 Maven 配置，确保映射文件被正确打包到 JAR 或 WAR 文件中

```xml
<!-- 在 pom.xml 中添加配置，确保 XML 文件被正确打包 -->
<build>
    <resources>
        <resource>
            <directory>src/main/java</directory>
            <includes>
                <include>**/*.xml</include>
            </includes>
        </resource>
        <resource>
            <directory>src/main/resources</directory>
        </resource>
    </resources>
</build>
```

#### 10.1.2 类型转换错误

**问题**：执行 SQL 语句时出现类型转换错误，如 `java.lang.ClassCastException`。

**解决方案**：
- 检查 Java 类型和数据库类型是否匹配
- 使用合适的类型处理器（TypeHandler）
- 检查映射文件中的 `parameterType` 和 `resultType` 是否正确

### 10.2 性能问题

#### 10.2.1 N+1 查询问题

**问题**：执行关联查询时出现 N+1 查询问题，导致性能下降。

**解决方案**：
- 使用联表查询替代嵌套查询
- 使用 `fetchType="eager"` 来立即加载关联数据
- 使用 `@Results` 注解来配置关联查询

```xml
<!-- 使用联表查询 -->
<select id="getUserWithRole" parameterType="int" resultMap="UserWithRoleMap">
    SELECT u.id, u.username, u.email, r.id as role_id, r.name as role_name
    FROM user u
    LEFT JOIN user_role ur ON u.id = ur.user_id
    LEFT JOIN role r ON ur.role_id = r.id
    WHERE u.id = #{id}
</select>

<resultMap id="UserWithRoleMap" type="User">
    <id property="id" column="id"/>
    <result property="username" column="username"/>
    <result property="email" column="email"/>
    <association property="role" javaType="Role">
        <id property="id" column="role_id"/>
        <result property="name" column="role_name"/>
    </association>
</resultMap>
```

#### 10.2.2 大量数据查询性能问题

**问题**：查询大量数据时出现性能问题，如内存溢出、响应时间过长等。

**解决方案**：
- 使用分页查询
- 使用流式查询
- 优化 SQL 语句，添加合适的索引
- 考虑使用缓存

### 10.3 事务问题

#### 10.3.1 事务不回滚

**问题**：执行事务操作时，出现异常但事务不回滚。

**解决方案**：
- 检查 `@Transactional` 注解的 `rollbackFor` 属性是否正确设置
- 检查异常是否被捕获并处理，导致事务不回滚
- 检查事务管理器的配置是否正确

```java
@Transactional(rollbackFor = Exception.class) // 确保所有异常都能触发回滚
public void createUser(User user) {
    try {
        userMapper.insertUser(user);
        // 其他操作
    } catch (Exception e) {
        // 记录日志
        throw e; // 重新抛出异常，确保事务回滚
    }
}
```

#### 10.3.2 事务传播问题

**问题**：嵌套事务时出现事务传播问题，如事务不生效、事务被意外提交等。

**解决方案**：
- 检查 `@Transactional` 注解的 `propagation` 属性是否正确设置
- 了解不同事务传播行为的含义，选择合适的传播行为
- 避免过度使用嵌套事务

### 10.4 其他常见问题

#### 10.4.1 循环引用问题

**问题**：在关联查询中出现循环引用问题，导致 JSON 序列化失败。

**解决方案**：
- 使用 `@JsonIgnore` 注解忽略循环引用的属性
- 使用 `@JsonManagedReference` 和 `@JsonBackReference` 注解来处理循环引用
- 使用 DTO（数据传输对象）来避免循环引用

#### 10.4.2 数据库连接泄漏

**问题**：应用运行一段时间后出现数据库连接泄漏，导致连接池耗尽。

**解决方案**：
- 确保 `SqlSession` 被正确关闭
- 使用 try-with-resources 语句来自动关闭资源
- 监控连接池的使用情况，及时发现和解决问题

```java
// 使用 try-with-resources 语句自动关闭 SqlSession
try (SqlSession sqlSession = sqlSessionFactory.openSession()) {
    UserMapper userMapper = sqlSession.getMapper(UserMapper.class);
    User user = userMapper.getUserById(1);
    // 使用 user
}
```

## 11. 结论和参考资料

### 11.1 结论

MyBatis 是一个功能强大、灵活易用的持久层框架，它提供了以下优势：

- **简单易用**：MyBatis 的 API 设计简洁明了，学习曲线平缓，容易上手。
- **灵活性高**：MyBatis 允许开发者直接编写 SQL 语句，具有很高的灵活性和可控性。
- **性能优秀**：MyBatis 提供了缓存机制、延迟加载、批量操作等功能，可以显著提高查询性能。
- **易于集成**：MyBatis 可以与 Spring、Spring Boot 等框架无缝集成，便于开发企业级应用。
- **丰富的功能**：MyBatis 提供了动态 SQL、类型处理器、插件机制等丰富的功能，可以满足各种复杂的业务需求。

MyBatis 适用于以下场景：

- 需要高度控制 SQL 语句的项目
- 对性能要求较高的项目
- 与 Spring、Spring Boot 等框架集成的项目
- 需要处理复杂数据映射的项目

### 11.2 参考资料

1. **MyBatis 官方文档**：
   - [MyBatis 3 官方文档](https://mybatis.org/mybatis-3/zh/index.html)
   - [MyBatis Spring 官方文档](https://mybatis.org/spring/zh/index.html)
   - [MyBatis Spring Boot Starter 官方文档](https://github.com/mybatis/spring-boot-starter)

2. **MyBatis 相关书籍**：
   - 《MyBatis 从入门到精通》
   - 《Spring Boot 实战》
   - 《Java 持久化技术》

3. **MyBatis 相关资源**：
   - [MyBatis GitHub 仓库](https://github.com/mybatis/mybatis-3)
   - [MyBatis Spring GitHub 仓库](https://github.com/mybatis/spring)
   - [MyBatis Spring Boot Starter GitHub 仓库](https://github.com/mybatis/spring-boot-starter)
   - [MyBatis-Plus 官方文档](https://baomidou.com/guide/)

4. **相关技术文档**：
   - [Spring 官方文档](https://spring.io/docs)
   - [Spring Boot 官方文档](https://spring.io/projects/spring-boot)
   - [MySQL 官方文档](https://dev.mysql.com/doc/)
   - [HikariCP 官方文档](https://github.com/brettwooldridge/HikariCP)

5. **学习资源**：
   - [MyBatis 官方示例](https://github.com/mybatis/mybatis-3/tree/master/src/test/java/org/apache/ibatis/submitted)
   - [MyBatis Spring Boot 示例](https://github.com/mybatis/spring-boot-starter/tree/master/mybatis-spring-boot-samples)
   - [B站 MyBatis 教学视频](https://www.bilibili.com/video/BV1NE411Q7Nx)
   - [慕课网 MyBatis 课程](https://www.imooc.com/search/?words=MyBatis)

通过本教程，我们详细介绍了 MyBatis 的基本概念、核心功能、高级特性以及与 Spring 框架的集成方法。希望本教程能够帮助你更好地理解和使用 MyBatis，提高你的开发效率和代码质量。
