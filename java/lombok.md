# Lombok 技术文档

## 1. 什么是 Lombok？

Lombok 是一个 Java 库，它通过注解的方式，在编译时期自动为 Java 类生成样板代码（boilerplate code），如 getter、setter、构造函数、equals、hashCode 和 toString 方法等。使用 Lombok 可以大大减少代码量，提高开发效率，使代码更加简洁易读。

## 2. Lombok 的优势

- **减少样板代码**：自动生成 getter、setter、构造函数等常用方法
- **提高代码可读性**：移除重复性代码，使核心业务逻辑更加清晰
- **减少维护成本**：当类的属性发生变化时，不需要手动更新相关方法
- **编译期处理**：在编译时期生成代码，不影响运行时性能
- **易于使用**：通过简单的注解即可实现复杂功能

## 3. Lombok 安装

### 3.1 Maven 依赖

在 Maven 项目的 `pom.xml` 文件中添加以下依赖：

```xml
<dependency>
    <groupId>org.projectlombok</groupId>
    <artifactId>lombok</artifactId>
    <version>1.18.30</version>
    <scope>provided</scope>
</dependency>
```

### 3.2 Gradle 依赖

在 Gradle 项目的 `build.gradle` 文件中添加以下依赖：

```groovy
dependencies {
    compileOnly 'org.projectlombok:lombok:1.18.30'
    annotationProcessor 'org.projectlombok:lombok:1.18.30'
}
```

### 3.3 IDE 插件安装

为了在 IDE 中获得更好的支持，需要安装相应的 Lombok 插件：

#### IntelliJ IDEA

1. 打开 Settings -> Plugins
2. 搜索 "Lombok Plugin"
3. 安装并重启 IDE

#### Eclipse

1. 打开 Help -> Eclipse Marketplace
2. 搜索 "Lombok"
3. 安装并重启 IDE

## 4. Lombok 常用注解

### 4.1 基础注解

#### @Getter / @Setter

自动生成字段的 getter 和 setter 方法。

```java
import lombok.Getter;
import lombok.Setter;

public class User {
    @Getter @Setter private String name;
    @Getter @Setter private int age;
}
```

等价于：

```java
public class User {
    private String name;
    private int age;
    
    public String getName() {
        return name;
    }
    
    public void setName(String name) {
        this.name = name;
    }
    
    public int getAge() {
        return age;
    }
    
    public void setAge(int age) {
        this.age = age;
    }
}
```

#### @ToString

自动生成 `toString()` 方法，包含类名和所有非静态字段。

```java
import lombok.ToString;

@ToString
public class User {
    private String name;
    private int age;
    private String email;
}
```

等价于：

```java
public class User {
    private String name;
    private int age;
    private String email;
    
    @Override
    public String toString() {
        return "User(name=" + name + ", age=" + age + ", email=" + email + ")";
    }
}
```

可以通过 `exclude` 参数排除特定字段：

```java
@ToString(exclude = {"email"})
public class User {
    private String name;
    private int age;
    private String email;
}
```

#### @EqualsAndHashCode

自动生成 `equals()` 和 `hashCode()` 方法。

```java
import lombok.EqualsAndHashCode;

@EqualsAndHashCode
public class User {
    private String name;
    private int age;
    private String email;
}
```

可以通过 `of` 参数指定用于比较的字段：

```java
@EqualsAndHashCode(of = {"name", "email"})
public class User {
    private String name;
    private int age;
    private String email;
}
```

#### @NoArgsConstructor, @RequiredArgsConstructor, @AllArgsConstructor

生成构造函数：
- `@NoArgsConstructor`：生成无参构造函数
- `@RequiredArgsConstructor`：为所有 final 或 @NonNull 字段生成构造函数
- `@AllArgsConstructor`：为所有字段生成构造函数

```java
import lombok.NoArgsConstructor;
import lombok.RequiredArgsConstructor;
import lombok.AllArgsConstructor;
import lombok.NonNull;

@NoArgsConstructor
@AllArgsConstructor
@RequiredArgsConstructor
public class User {
    @NonNull private String name;
    private int age;
    private String email;
}
```

等价于：

```java
public class User {
    private String name;
    private int age;
    private String email;
    
    public User() {
    }
    
    public User(String name) {
        if (name == null) {
            throw new NullPointerException("name is marked non-null but is null");
        }
        this.name = name;
    }
    
    public User(String name, int age, String email) {
        if (name == null) {
            throw new NullPointerException("name is marked non-null but is null");
        }
        this.name = name;
        this.age = age;
        this.email = email;
    }
}
```

### 4.2 组合注解

#### @Data

组合了 @Getter、@Setter、@ToString、@EqualsAndHashCode 和 @RequiredArgsConstructor 注解的功能。

```java
import lombok.Data;

@Data
public class User {
    private String name;
    private int age;
    private String email;
}
```

等价于：

```java
public class User {
    private String name;
    private int age;
    private String email;
    
    // 自动生成 getter 和 setter 方法
    // 自动生成 toString() 方法
    // 自动生成 equals() 和 hashCode() 方法
    // 自动生成为 final 或 @NonNull 字段的构造函数
}
```

#### @Value

@Value 注解用于创建不可变类，它组合了 @Getter、@ToString、@EqualsAndHashCode、@RequiredArgsConstructor 注解的功能，并将所有字段默认设为 final。

```java
import lombok.Value;

@Value
public class User {
    String name;
    int age;
    String email;
}
```

等价于：

```java
public final class User {
    private final String name;
    private final int age;
    private final String email;
    
    // 自动生成 getter 方法（没有 setter 方法）
    // 自动生成 toString() 方法
    // 自动生成 equals() 和 hashCode() 方法
    // 自动生成包含所有字段的构造函数
}
```

#### @Builder

实现构建者模式，允许通过链式调用创建对象。

```java
import lombok.Builder;

@Builder
public class User {
    private String name;
    private int age;
    private String email;
}
```

使用方式：

```java
User user = User.builder()
    .name("张三")
    .age(25)
    .email("zhangsan@example.com")
    .build();
```

可以与 @Data 等注解组合使用：

```java
import lombok.Data;
import lombok.Builder;

@Data
@Builder
public class User {
    private String name;
    private int age;
    private String email;
}
```

### 4.3 其他常用注解

#### @NonNull

用于字段或方法参数，生成非空检查代码。

```java
import lombok.NonNull;

public class User {
    private String name;
    
    public void setName(@NonNull String name) {
        this.name = name;
    }
}
```

等价于：

```java
public class User {
    private String name;
    
    public void setName(String name) {
        if (name == null) {
            throw new NullPointerException("name is marked non-null but is null");
        }
        this.name = name;
    }
}
```

#### @Cleanup

自动关闭资源，相当于 try-with-resources 语句。

```java
import lombok.Cleanup;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

public class FileCopyExample {
    public static void copyFile(String source, String destination) throws IOException {
        @Cleanup FileInputStream in = new FileInputStream(source);
        @Cleanup FileOutputStream out = new FileOutputStream(destination);
        
        byte[] buffer = new byte[1024];
        int len;
        while ((len = in.read(buffer)) != -1) {
            out.write(buffer, 0, len);
        }
    }
}
```

等价于：

```java
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

public class FileCopyExample {
    public static void copyFile(String source, String destination) throws IOException {
        try (FileInputStream in = new FileInputStream(source);
             FileOutputStream out = new FileOutputStream(destination)) {
            
            byte[] buffer = new byte[1024];
            int len;
            while ((len = in.read(buffer)) != -1) {
                out.write(buffer, 0, len);
            }
        }
    }
}
```

#### @SneakyThrows

自动捕获并包装异常，避免在方法签名中声明受检异常。

```java
import lombok.SneakyThrows;

public class SneakyThrowsExample {
    @SneakyThrows(Exception.class)
    public void doSomething() {
        Thread.sleep(1000); // 不需要显式捕获 InterruptedException
    }
}
```

等价于：

```java
public class SneakyThrowsExample {
    public void doSomething() {
        try {
            Thread.sleep(1000);
        } catch (Exception e) {
            throw Lombok.sneakyThrow(e);
        }
    }
}
```

#### @Log

为类生成日志记录器。Lombok 支持多种日志框架，包括：
- @Log：java.util.logging
- @Slf4j：SLF4J
- @Log4j：Log4j
- @Log4j2：Log4j2
- @CommonsLog：Apache Commons Logging

```java
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class LogExample {
    public void doSomething() {
        log.info("正在执行操作...");
        // 业务逻辑
        log.debug("操作执行完成");
    }
}
```

等价于：

```java
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class LogExample {
    private static final Logger log = LoggerFactory.getLogger(LogExample.class);
    
    public void doSomething() {
        log.info("正在执行操作...");
        // 业务逻辑
        log.debug("操作执行完成");
    }
}
```

## 5. Lombok 高级特性

### 5.1 自定义配置

Lombok 提供了配置文件 `lombok.config`，可以在项目根目录创建该文件来自定义 Lombok 的行为。

示例 `lombok.config` 文件：

```properties
# 启用 Lombok 的所有实验性特性
lombok.experimental=true

# 设置生成的 getter/setter 方法的前缀
lombok.accessors.prefix=_

# 设置生成的 toString 方法的格式
lombok.toString.callSuper=call

# 设置生成的构造函数的访问级别
lombok.anyConstructor.accessLevel=protected

# 禁用特定类的 Lombok 注解
lombok.addLombokGeneratedAnnotation=true
```

### 5.2 访问器（Accessors）配置

使用 @Accessors 注解可以自定义 getter/setter 方法的生成规则。

```java
import lombok.Getter;
import lombok.Setter;
import lombok.experimental.Accessors;

@Getter
@Setter
@Accessors(fluent = true, chain = true, prefix = "_")
public class User {
    private String _name;
    private int _age;
}
```

使用方式：

```java
User user = new User();
user.name("张三").age(25); // fluent 和 chain 风格
String name = user.name(); // fluent 风格的 getter
```

### 5.3 扩展方法（Extension Methods）

Lombok 提供了 @ExtensionMethod 注解，允许为现有类添加扩展方法（类似于 Kotlin 的扩展函数）。

```java
import lombok.experimental.ExtensionMethod;

@ExtensionMethod({StringExtensions.class, NumberExtensions.class})
public class ExtensionMethodExample {
    public void example() {
        String str = "hello";
        boolean isEmpty = str.isEmpty(); // 标准方法
        String upper = str.toUpperCase(); // 标准方法
        String camelCase = str.toCamelCase(); // 扩展方法
        
        Integer num = 10;
        boolean isEven = num.isEven(); // 扩展方法
    }
}

class StringExtensions {
    public static String toCamelCase(String self) {
        if (self == null || self.isEmpty()) {
            return self;
        }
        String[] parts = self.split(" ");
        StringBuilder result = new StringBuilder(parts[0].toLowerCase());
        for (int i = 1; i < parts.length; i++) {
            result.append(parts[i].substring(0, 1).toUpperCase()).append(parts[i].substring(1).toLowerCase());
        }
        return result.toString();
    }
}

class NumberExtensions {
    public static boolean isEven(Integer self) {
        return self != null && self % 2 == 0;
    }
}
```

### 5.4 同步方法（Synchronized）

使用 @Synchronized 注解可以为方法添加同步锁，比使用 synchronized 关键字更加安全。

```java
import lombok.Synchronized;

public class SynchronizedExample {
    private final Object readLock = new Object();
    
    @Synchronized
    public void write() {
        // 写操作，使用 this 作为锁
    }
    
    @Synchronized("readLock")
    public void read() {
        // 读操作，使用 readLock 作为锁
    }
}
```

等价于：

```java
public class SynchronizedExample {
    private final Object readLock = new Object();
    private final Object $lock = new Object[0]; // Lombok 生成的锁对象
    
    public void write() {
        synchronized ($lock) {
            // 写操作
        }
    }
    
    public void read() {
        synchronized (readLock) {
            // 读操作
        }
    }
}
```

## 6. Lombok 与 IDE 的集成

### 6.1 IntelliJ IDEA

1. 安装 Lombok 插件：
   - 打开 Settings -> Plugins
   - 搜索 "Lombok Plugin"
   - 点击 "Install" 安装
   - 重启 IDE

2. 启用注解处理：
   - 打开 Settings -> Build, Execution, Deployment -> Compiler -> Annotation Processors
   - 勾选 "Enable annotation processing"

### 6.2 Eclipse

1. 下载 Lombok JAR 文件：
   - 从 [Lombok 官网](https://projectlombok.org/download) 下载最新版本的 JAR 文件

2. 安装 Lombok：
   - 双击下载的 JAR 文件
   - 在弹出的窗口中，选择 Eclipse 的安装目录
   - 点击 "Install / Update"
   - 重启 Eclipse

3. 启用注解处理：
   - 打开 Preferences -> Java -> Compiler -> Annotation Processing
   - 勾选 "Enable annotation processing"

### 6.3 NetBeans

1. 安装 Lombok 插件：
   - 打开 Tools -> Plugins
   - 点击 "Downloaded" 标签
   - 点击 "Add Plugins..."
   - 选择下载的 Lombok NBM 文件
   - 点击 "Install"
   - 重启 NetBeans

## 7. Lombok 最佳实践

### 7.1 何时使用 Lombok

- **数据传输对象（DTO）**：使用 @Data 或 @Value 注解可以快速创建 DTO 类
- **实体类（Entity）**：使用 @Data、@Builder 等注解可以简化实体类的编写
- **工具类**：使用 @UtilityClass 注解可以创建工具类
- **日志记录**：使用 @Slf4j 等注解可以方便地添加日志功能

### 7.2 何时避免使用 Lombok

- **需要自定义 getter/setter 逻辑**：当 getter/setter 方法需要包含业务逻辑时，不建议使用 Lombok
- **与其他注解处理器冲突**：当项目中使用了其他可能与 Lombok 冲突的注解处理器时，需要谨慎使用
- **团队成员不熟悉 Lombok**：当团队成员不熟悉 Lombok 时，可能会导致代码维护困难
- **项目对编译速度有严格要求**：使用 Lombok 会增加编译时间

### 7.3 最佳实践建议

1. **统一团队规范**：团队成员应共同约定 Lombok 的使用规范，包括哪些注解可以使用，如何命名等

2. **合理使用组合注解**：
   - 使用 @Data 时要注意，它会生成 equals() 和 hashCode() 方法，可能会影响性能
   - 对于实体类，可以使用 @Getter、@Setter、@ToString 等单独注解，而不是 @Data

3. **注意继承关系**：
   - 使用 @ToString(callSuper = true) 可以在 toString() 方法中包含父类的字段
   - 使用 @EqualsAndHashCode(callSuper = true) 可以在 equals() 和 hashCode() 方法中包含父类的字段

4. **避免过度使用 @Builder**：
   - @Builder 会生成大量代码，增加编译时间
   - 对于简单的类，可以使用构造函数或 setter 方法代替

5. **使用 @Value 创建不可变对象**：
   - 不可变对象更加安全，不容易出现并发问题
   - 可以使用 @With 注解为不可变对象提供属性更新方法

6. **注意 @NonNull 注解的使用**：
   - @NonNull 会生成非空检查代码，可能会影响性能
   - 只对确实需要非空检查的字段使用 @NonNull

## 8. Lombok 与其他框架的集成

### 8.1 Spring Boot

Lombok 与 Spring Boot 集成非常简单，只需在 pom.xml 中添加 Lombok 依赖即可。

```xml
<dependencies>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    <dependency>
        <groupId>org.projectlombok</groupId>
        <artifactId>lombok</artifactId>
        <optional>true</optional>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-test</artifactId>
        <scope>test</scope>
    </dependency>
</dependencies>
```

使用示例：

```java
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.AllArgsConstructor;
import lombok.Builder;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String name;
    private int age;
    private String email;
}
```

### 8.2 Hibernate/JPA

Lombok 与 Hibernate/JPA 集成时，需要注意以下几点：

1. **实体类注解**：使用 @Entity、@Table 等 JPA 注解的同时，可以使用 @Data、@Builder 等 Lombok 注解

2. **字段映射**：使用 @Column、@JoinColumn 等注解进行字段映射

3. **关联关系**：在处理关联关系时，需要注意避免循环引用导致的 StackOverflowError

```java
import lombok.Data;
import lombok.ToString;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.OneToMany;
import java.util.List;

@Entity
@Data
@ToString(exclude = {"orders"})
public class Customer {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String name;
    private String email;
    
    @OneToMany(mappedBy = "customer")
    private List<Order> orders;
}

@Entity
@Data
@ToString(exclude = {"customer"})
public class Order {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String product;
    private int quantity;
    
    @ManyToOne
    private Customer customer;
}
```

### 8.3 Jackson/Gson

Lombok 生成的 getter/setter 方法可以被 Jackson、Gson 等 JSON 序列化框架正常识别和使用。

```java
import lombok.Data;
import com.fasterxml.jackson.annotation.JsonProperty;

@Data
public class UserDto {
    @JsonProperty("user_name")
    private String name;
    private int age;
    private String email;
}
```

## 9. Lombok 常见问题及解决方案

### 9.1 IDE 不识别 Lombok 生成的方法

**问题**：在 IDE 中，使用 Lombok 注解的类，其生成的方法（如 getter、setter）无法被识别，导致编译错误或代码补全失效。

**解决方案**：
1. 确保已安装 Lombok 插件
2. 启用 IDE 的注解处理功能
3. 清理和重建项目
4. 重启 IDE

### 9.2 Lombok 与其他注解处理器冲突

**问题**：当项目中同时使用 Lombok 和其他注解处理器（如 MapStruct、QueryDSL 等）时，可能会出现冲突。

**解决方案**：
1. 确保所有注解处理器的版本兼容
2. 在构建工具配置中正确配置注解处理器的执行顺序
3. 避免在同一个类上使用可能冲突的注解

### 9.3 序列化/反序列化问题

**问题**：使用 Lombok 注解的类在序列化或反序列化时出现问题。

**解决方案**：
1. 确保类有默认构造函数（可以使用 @NoArgsConstructor）
2. 对于不可变类，可以使用 @Jacksonized（与 @Builder 配合使用）
3. 自定义序列化/反序列化逻辑时，需要手动编写相关方法

### 9.4 继承关系中的问题

**问题**：在继承关系中，使用 Lombok 注解可能会导致 equals()、hashCode() 或 toString() 方法不正确。

**解决方案**：
1. 使用 @ToString(callSuper = true) 包含父类字段
2. 使用 @EqualsAndHashCode(callSuper = true) 包含父类字段
3. 对于复杂的继承关系，考虑手动编写这些方法

## 10. Lombok 替代方案

虽然 Lombok 是一个非常流行的库，但也有一些替代方案：

### 10.1 Java 14+ 记录类（Records）

Java 14 引入了记录类（Records），它是一种特殊的类，用于简洁地声明数据载体类。

```java
public record User(String name, int age, String email) {
    // 自动生成构造函数、getter、equals、hashCode 和 toString 方法
}
```

### 10.2 Kotlin 数据类（Data Classes）

如果你正在考虑使用 Kotlin，那么它的数���类是一个很好的选择：

```kotlin
data class User(val name: String, val age: Int, val email: String)
```

### 10.3 MapStruct

MapStruct 是一个代码生成器，用于简化 Java bean 之间的映射：

```java
@Mapper
public interface UserMapper {
    UserMapper INSTANCE = Mappers.getMapper(UserMapper.class);
    
    @Mapping(source = "name", target = "userName")
    UserDto userToUserDto(User user);
    
    List<UserDto> usersToUserDtos(List<User> users);
}
```

## 11. 总结

Lombok 是一个强大的 Java 库，它可以通过注解的方式自动生成样板代码，大大减少代码量，提高开发效率。使用 Lombok 可以使代码更加简洁易读，减少维护成本。

然而，使用 Lombok 也需要注意一些问题，如与其他注解处理器的冲突、IDE 支持问题等。在使用 Lombok 时，应该遵循最佳实践，合理使用各种注解，避免过度使用。

对于新项目，可以考虑使用 Java 14+ 的记录类或 Kotlin 数据类作为替代方案，它们提供了类似的功能，但不需要额外的依赖。

总之，Lombok 是一个非常有用的工具，它可以帮助开发者编写更加简洁、高效的 Java 代码，但需要在合适的场景下合理使用。
