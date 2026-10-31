# Jackson 技术文档

## 1. 什么是 Jackson？

Jackson 是一个功能强大的 Java 库，用于处理 JSON 格式的数据。它提供了将 Java 对象与 JSON 数据进行相互转换（序列化和反序列化）的功能，是 Java 生态系统中最流行的 JSON 处理库之一。Jackson 具有高性能、灵活性和广泛的功能集，使其成为开发人员处理 JSON 数据的首选工具。

## 2. Jackson 的主要模块

Jackson 采用模块化设计，由多个核心模块组成，每个模块负责不同的功能：

- **jackson-core**：核心模块，提供底层的 JSON 解析和生成功能
- **jackson-databind**：数据绑定模块，提供 Java 对象与 JSON 之间的转换功能
- **jackson-annotations**：注解模块，提供用于自定义序列化和反序列化行为的注解
- **jackson-dataformat-xml**：XML 格式支持模块
- **jackson-dataformat-yaml**：YAML 格式支持模块
- **jackson-dataformat-csv**：CSV 格式支持模块
- **jackson-module-jaxb-annotations**：JAXB 注解支持模块
- **jackson-module-kotlin**：Kotlin 语言支持模块

## 3. Jackson 的优势

- **高性能**：Jackson 是目前性能最高的 JSON 处理库之一
- **功能丰富**：提供了广泛的功能，包括数据绑定、树模型、流式 API 等
- **灵活性**：支持多种数据格式和自定义序列化/反序列化逻辑
- **广泛的社区支持**：是 Java 生态系统中最流行的 JSON 处理库之一
- **与 Spring 框架的良好集成**：Spring 框架默认使用 Jackson 处理 JSON 数据

## 4. Jackson 安装

### 4.1 Maven 依赖

在 Maven 项目的 `pom.xml` 文件中添加以下依赖：

```xml
<!-- 核心模块 -->
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-core</artifactId>
    <version>2.15.2</version>
</dependency>

<!-- 数据绑定模块 -->
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-databind</artifactId>
    <version>2.15.2</version>
</dependency>

<!-- 注解模块 -->
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-annotations</artifactId>
    <version>2.15.2</version>
</dependency>

<!-- 可选：XML 格式支持 -->
<dependency>
    <groupId>com.fasterxml.jackson.dataformat</groupId>
    <artifactId>jackson-dataformat-xml</artifactId>
    <version>2.15.2</version>
</dependency>

<!-- 可选：YAML 格式支持 -->
<dependency>
    <groupId>com.fasterxml.jackson.dataformat</groupId>
    <artifactId>jackson-dataformat-yaml</artifactId>
    <version>2.15.2</version>
</dependency>

<!-- 可选：JAXB 注解支持 -->
<dependency>
    <groupId>com.fasterxml.jackson.module</groupId>
    <artifactId>jackson-module-jaxb-annotations</artifactId>
    <version>2.15.2</version>
</dependency>
```

### 4.2 Gradle 依赖

在 Gradle 项目的 `build.gradle` 文件中添加以下依赖：

```groovy
dependencies {
    // 核心模块
    implementation 'com.fasterxml.jackson.core:jackson-core:2.15.2'
    
    // 数据绑定模块
    implementation 'com.fasterxml.jackson.core:jackson-databind:2.15.2'
    
    // 注解模块
    implementation 'com.fasterxml.jackson.core:jackson-annotations:2.15.2'
    
    // 可选：XML 格式支持
    implementation 'com.fasterxml.jackson.dataformat:jackson-dataformat-xml:2.15.2'
    
    // 可选：YAML 格式支持
    implementation 'com.fasterxml.jackson.dataformat:jackson-dataformat-yaml:2.15.2'
    
    // 可选：JAXB 注解支持
    implementation 'com.fasterxml.jackson.module:jackson-module-jaxb-annotations:2.15.2'
}
```

## 5. Jackson 核心功能

### 5.1 数据绑定 (ObjectMapper)

Jackson 的核心功能是通过 `ObjectMapper` 类实现的，它提供了 Java 对象与 JSON 数据之间的相互转换功能。

#### 5.1.1 基本序列化（Java 对象转 JSON）

```java
import com.fasterxml.jackson.databind.ObjectMapper;

public class SerializationExample {
    public static void main(String[] args) throws Exception {
        // 创建 ObjectMapper 实例
        ObjectMapper objectMapper = new ObjectMapper();
        
        // 创建 Java 对象
        User user = new User();
        user.setId(1);
        user.setName("张三");
        user.setEmail("zhangsan@example.com");
        
        // 序列化 Java 对象为 JSON 字符串
        String json = objectMapper.writeValueAsString(user);
        System.out.println(json);
        // 输出: {"id":1,"name":"张三","email":"zhangsan@example.com"}
        
        // 序列化 Java 对象为 JSON 文件
        objectMapper.writeValue(new File("user.json"), user);
    }
    
    static class User {
        private int id;
        private String name;
        private String email;
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
        public String getEmail() { return email; }
        public void setEmail(String email) { this.email = email; }
    }
}
```

#### 5.1.2 基本反序列化（JSON 转 Java 对象）

```java
import com.fasterxml.jackson.databind.ObjectMapper;

public class DeserializationExample {
    public static void main(String[] args) throws Exception {
        // 创建 ObjectMapper 实例
        ObjectMapper objectMapper = new ObjectMapper();
        
        // JSON 字符串
        String json = "{\"id\":1,\"name\":\"张三\",\"email\":\"zhangsan@example.com\"}";
        
        // 反序列化 JSON 字符串为 Java 对象
        User user = objectMapper.readValue(json, User.class);
        System.out.println(user.getName()); // 输出: 张三
        
        // 从 JSON 文件反序列化为 Java 对象
        User userFromFile = objectMapper.readValue(new File("user.json"), User.class);
        System.out.println(userFromFile.getName()); // 输出: 张三
    }
    
    static class User {
        private int id;
        private String name;
        private String email;
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
        public String getEmail() { return email; }
        public void setEmail(String email) { this.email = email; }
    }
}
```

#### 5.1.3 集合类型的序列化和反序列化

```java
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.util.List;
import java.util.Map;

public class CollectionExample {
    public static void main(String[] args) throws Exception {
        ObjectMapper objectMapper = new ObjectMapper();
        
        // 序列化 List
        List<User> users = List.of(
            new User(1, "张三", "zhangsan@example.com"),
            new User(2, "李四", "lisi@example.com")
        );
        String jsonUsers = objectMapper.writeValueAsString(users);
        System.out.println(jsonUsers);
        // 输出: [{"id":1,"name":"张三","email":"zhangsan@example.com"},{"id":2,"name":"李四","email":"lisi@example.com"}]
        
        // 反序列化 JSON 数组为 List
        List<User> deserializedUsers = objectMapper.readValue(
            jsonUsers, 
            new TypeReference<List<User>>() {}
        );
        System.out.println(deserializedUsers.size()); // 输出: 2
        
        // 序列化 Map
        Map<String, Object> userMap = Map.of(
            "id", 1,
            "name", "张三",
            "email", "zhangsan@example.com"
        );
        String jsonMap = objectMapper.writeValueAsString(userMap);
        System.out.println(jsonMap);
        // 输出: {"name":"张三","email":"zhangsan@example.com","id":1}
        
        // 反序列化 JSON 对象为 Map
        Map<String, Object> deserializedMap = objectMapper.readValue(
            jsonMap, 
            new TypeReference<Map<String, Object>>() {}
        );
        System.out.println(deserializedMap.get("name")); // 输出: 张三
    }
    
    static class User {
        private int id;
        private String name;
        private String email;
        
        public User(int id, String name, String email) {
            this.id = id;
            this.name = name;
            this.email = email;
        }
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
        public String getEmail() { return email; }
        public void setEmail(String email) { this.email = email; }
    }
}
```

### 5.2 树模型 (JsonNode)

Jackson 提供了树模型 API，使用 `JsonNode` 类表示 JSON 数据的树形结构，允许以灵活的方式访问和修改 JSON 数据。

```java
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ObjectNode;

public class TreeModelExample {
    public static void main(String[] args) throws Exception {
        ObjectMapper objectMapper = new ObjectMapper();
        
        // 解析 JSON 字符串为 JsonNode
        String json = "{\"id\":1,\"name\":\"张三\",\"email\":\"zhangsan@example.com\",\"address\":{\"city\":\"北京\",\"street\":\"朝阳区\"}}";
        JsonNode rootNode = objectMapper.readTree(json);
        
        // 访问 JSON 字段
        int id = rootNode.get("id").asInt();
        String name = rootNode.get("name").asText();
        String email = rootNode.get("email").asText();
        String city = rootNode.get("address").get("city").asText();
        
        System.out.println("ID: " + id); // 输出: ID: 1
        System.out.println("Name: " + name); // 输出: Name: 张三
        System.out.println("Email: " + email); // 输出: Email: zhangsan@example.com
        System.out.println("City: " + city); // 输出: City: 北京
        
        // 修改 JSON 字段
        ((ObjectNode) rootNode).put("name", "李四");
        ((ObjectNode) rootNode).put("age", 25);
        
        // 添加新的嵌套对象
        ObjectNode contactNode = ((ObjectNode) rootNode).putObject("contact");
        contactNode.put("phone", "13800138000");
        contactNode.put("wechat", "lisi123");
        
        // 删除字段
        ((ObjectNode) rootNode).remove("email");
        
        // 转换回 JSON 字符串
        String updatedJson = objectMapper.writeValueAsString(rootNode);
        System.out.println(updatedJson);
        // 输出: {"id":1,"name":"李四","address":{"city":"北京","street":"朝阳区"},"age":25,"contact":{"phone":"13800138000","wechat":"lisi123"}}
    }
}
```

### 5.3 流式 API (JsonParser/JsonGenerator)

Jackson 提供了流式 API，通过 `JsonParser` 和 `JsonGenerator` 类实现，允许以流式方式处理 JSON 数据，适用于处理大型 JSON 数据或需要高效处理的场景。

#### 5.3.1 使用 JsonGenerator 生成 JSON

```java
import com.fasterxml.jackson.core.JsonFactory;
import com.fasterxml.jackson.core.JsonGenerator;
import java.io.File;
import java.io.FileWriter;

public class JsonGeneratorExample {
    public static void main(String[] args) throws Exception {
        JsonFactory factory = new JsonFactory();
        
        // 使用 FileWriter 输出 JSON
        try (JsonGenerator generator = factory.createGenerator(new FileWriter("users_stream.json"))) {
            generator.writeStartObject(); // 开始 JSON 对象 {"
            
            generator.writeNumberField("id", 1);
            generator.writeStringField("name", "张三");
            generator.writeStringField("email", "zhangsan@example.com");
            
            // 嵌套对象
            generator.writeFieldName("address");
            generator.writeStartObject(); // 开始嵌套对象 {"
            generator.writeStringField("city", "北京");
            generator.writeStringField("street", "朝阳区");
            generator.writeEndObject(); // 结束嵌套对象 }
            
            // 数组
            generator.writeFieldName("hobbies");
            generator.writeStartArray(); // 开始数组 ["
            generator.writeString("读书");
            generator.writeString("旅行");
            generator.writeString("运动");
            generator.writeEndArray(); // 结束数组 ]"
            
            generator.writeEndObject(); // 结束 JSON 对象 }"
        }
        
        System.out.println("JSON 文件生成成功!");
    }
}
```

#### 5.3.2 使用 JsonParser 解析 JSON

```java
import com.fasterxml.jackson.core.JsonFactory;
import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.JsonToken;
import java.io.File;

public class JsonParserExample {
    public static void main(String[] args) throws Exception {
        JsonFactory factory = new JsonFactory();
        
        try (JsonParser parser = factory.createParser(new File("users_stream.json"))) {
            while (!parser.isClosed()) {
                JsonToken token = parser.nextToken();
                
                if (token == null) {
                    break;
                }
                
                if (JsonToken.FIELD_NAME.equals(token)) {
                    String fieldName = parser.getCurrentName();
                    token = parser.nextToken();
                    
                    if ("id".equals(fieldName)) {
                        System.out.println("ID: " + parser.getValueAsInt());
                    } else if ("name".equals(fieldName)) {
                        System.out.println("Name: " + parser.getValueAsString());
                    } else if ("email".equals(fieldName)) {
                        System.out.println("Email: " + parser.getValueAsString());
                    } else if ("address".equals(fieldName)) {
                        while (!JsonToken.END_OBJECT.equals(token = parser.nextToken())) {
                            if (JsonToken.FIELD_NAME.equals(token)) {
                                String addressField = parser.getCurrentName();
                                token = parser.nextToken();
                                if ("city".equals(addressField)) {
                                    System.out.println("City: " + parser.getValueAsString());
                                } else if ("street".equals(addressField)) {
                                    System.out.println("Street: " + parser.getValueAsString());
                                }
                            }
                        }
                    } else if ("hobbies".equals(fieldName)) {
                        System.out.print("Hobbies: ");
                        while (!JsonToken.END_ARRAY.equals(token = parser.nextToken())) {
                            if (JsonToken.VALUE_STRING.equals(token)) {
                                System.out.print(parser.getValueAsString() + ", ");
                            }
                        }
                        System.out.println();
                    }
                }
            }
        }
    }
}
```

## 6. Jackson 注解

Jackson 提供了丰富的注解，用于自定义序列化和反序列化行为。以下是一些常用的注解：

### 6.1 基本注解

#### @JsonProperty

用于指定 JSON 字段的名称，解决 Java 属性名与 JSON 字段名不匹配的问题。

```java
import com.fasterxml.jackson.annotation.JsonProperty;

public class User {
    @JsonProperty("user_id")
    private int id;
    
    @JsonProperty("user_name")
    private String name;
    
    @JsonProperty("user_email")
    private String email;
    
    // getters and setters
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
}
```

序列化后输出：
```json
{"user_id":1,"user_name":"张三","user_email":"zhangsan@example.com"}
```

#### @JsonIgnore

用于忽略指定的字段，使其在序列化和反序列化过程中不被处理。

```java
import com.fasterxml.jackson.annotation.JsonIgnore;

public class User {
    private int id;
    private String name;
    
    @JsonIgnore
    private String password; // 密码字段将被忽略
    
    private String email;
    
    // getters and setters
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public String getPassword() { return password; }
    public void setPassword(String password) { this.password = password; }
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
}
```

序列化后输出：
```json
{"id":1,"name":"张三","email":"zhangsan@example.com"}
```

#### @JsonIgnoreProperties

用于忽略多个字段，或忽略未知的 JSON 字段。

```java
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties({"password", "createdAt"}) // 忽略指定字段
// 或 @JsonIgnoreProperties(ignoreUnknown = true) // 忽略未知字段
public class User {
    private int id;
    private String name;
    private String password;
    private String email;
    private String createdAt;
    
    // getters and setters
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public String getPassword() { return password; }
    public void setPassword(String password) { this.password = password; }
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
    public String getCreatedAt() { return createdAt; }
    public void setCreatedAt(String createdAt) { this.createdAt = createdAt; }
}
```

#### @JsonInclude

用于指定何时包含字段，例如排除 null 值或空值。

```java
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL) // 排除 null 值
// 或 @JsonInclude(JsonInclude.Include.NON_EMPTY) // 排除 null 和空值
public class User {
    private int id;
    private String name;
    private String email; // 可能为 null
    private String phone; // 可能为 null
    
    // getters and setters
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
    public String getPhone() { return phone; }
    public void setPhone(String phone) { this.phone = phone; }
}
```

如果 email 和 phone 为 null，序列化后输出：
```json
{"id":1,"name":"张三"}
```

### 6.2 高级注解

#### @JsonFormat

用于格式化日期、时间和数字等类型的字段。

```java
import com.fasterxml.jackson.annotation.JsonFormat;
import java.util.Date;

public class Event {
    private String title;
    
    @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd HH:mm:ss", timezone = "GMT+8")
    private Date startTime;
    
    @JsonFormat(shape = JsonFormat.Shape.NUMBER)
    private Date endTime;
    
    @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "#,###.00")
    private double price;
    
    // getters and setters
    public String getTitle() { return title; }
    public void setTitle(String title) { this.title = title; }
    public Date getStartTime() { return startTime; }
    public void setStartTime(Date startTime) { this.startTime = startTime; }
    public Date getEndTime() { return endTime; }
    public void setEndTime(Date endTime) { this.endTime = endTime; }
    public double getPrice() { return price; }
    public void setPrice(double price) { this.price = price; }
}
```

序列化后输出：
```json
{"title":"会议","startTime":"2023-10-01 10:00:00","endTime":1696171200000,"price":"1,000.00"}
```

#### @JsonTypeInfo 和 @JsonSubTypes

用于处理多态类型的序列化和反序列化。

```java
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;

@JsonTypeInfo(
    use = JsonTypeInfo.Id.NAME,
    include = JsonTypeInfo.As.PROPERTY,
    property = "type"
)
@JsonSubTypes({
    @JsonSubTypes.Type(value = Car.class, name = "car"),
    @JsonSubTypes.Type(value = Bike.class, name = "bike")
})
public abstract class Vehicle {
    private String brand;
    private String model;
    
    // getters and setters
    public String getBrand() { return brand; }
    public void setBrand(String brand) { this.brand = brand; }
    public String getModel() { return model; }
    public void setModel(String model) { this.model = model; }
}

class Car extends Vehicle {
    private int doors;
    
    // getters and setters
    public int getDoors() { return doors; }
    public void setDoors(int doors) { this.doors = doors; }
}

class Bike extends Vehicle {
    private boolean hasBasket;
    
    // getters and setters
    public boolean isHasBasket() { return hasBasket; }
    public void setHasBasket(boolean hasBasket) { this.hasBasket = hasBasket; }
}
```

序列化后输出：
```json
{"type":"car","brand":"Toyota","model":"Camry","doors":4}
{"type":"bike","brand":"Giant","model":"Escape","hasBasket":true}
```

#### @JsonSerialize 和 @JsonDeserialize

用于自定义字段的序列化和反序列化逻辑。

```java
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class User {
    private int id;
    private String name;
    
    @JsonSerialize(using = PasswordSerializer.class)
    @JsonDeserialize(using = PasswordDeserializer.class)
    private String password;
    
    // getters and setters
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public String getPassword() { return password; }
    public void setPassword(String password) { this.password = password; }
}

// 自定义序列化器
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;
import java.io.IOException;

public class PasswordSerializer extends JsonSerializer<String> {
    @Override
    public void serialize(String value, JsonGenerator gen, SerializerProvider serializers) throws IOException {
        // 序列化时隐藏密码，只显示长度
        gen.writeString("******** (长度: " + value.length() + ")");
    }
}

// 自定义反序列化器
import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import java.io.IOException;

public class PasswordDeserializer extends JsonDeserializer<String> {
    @Override
    public String deserialize(JsonParser p, DeserializationContext ctxt) throws IOException {
        String value = p.getValueAsString();
        // 实际应用中可能需要进行密码验证或加密
        return value;
    }
}
```

#### @JsonView

用于控制不同场景下字段的可见性，实现数据视图功能。

```java
import com.fasterxml.jackson.annotation.JsonView;

public class User {
    public static class BasicView {};
    public static class DetailView extends BasicView {};
    
    @JsonView(BasicView.class)
    private int id;
    
    @JsonView(BasicView.class)
    private String name;
    
    @JsonView(DetailView.class)
    private String email;
    
    @JsonView(DetailView.class)
    private String phone;
    
    // getters and setters
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
    public String getPhone() { return phone; }
    public void setPhone(String phone) { this.phone = phone; }
}

// 使用示例
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;

public class JsonViewExample {
    public static void main(String[] args) throws Exception {
        ObjectMapper objectMapper = new ObjectMapper();
        objectMapper.enable(SerializationFeature.INDENT_OUTPUT);
        
        User user = new User();
        user.setId(1);
        user.setName("张三");
        user.setEmail("zhangsan@example.com");
        user.setPhone("13800138000");
        
        // 使用 BasicView 序列化
        String basicJson = objectMapper
            .writerWithView(User.BasicView.class)
            .writeValueAsString(user);
        System.out.println("Basic View:");
        System.out.println(basicJson);
        // 输出只包含 id 和 name 字段
        
        // 使用 DetailView 序列化
        String detailJson = objectMapper
            .writerWithView(User.DetailView.class)
            .writeValueAsString(user);
        System.out.println("\nDetail View:");
        System.out.println(detailJson);
        // 输出包含所有字段
    }
}
```

### 6.3 构造函数和工厂方法注解

#### @JsonCreator 和 @JsonProperty

用于自定义反序列化时使用的构造函数或工厂方法。

```java
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

public class User {
    private int id;
    private String name;
    private String email;
    
    // 自定义构造函数
    @JsonCreator
    public User(
        @JsonProperty("id") int id,
        @JsonProperty("name") String name,
        @JsonProperty("email") String email) {
        this.id = id;
        this.name = name;
        this.email = email;
    }
    
    // 或使用静态工厂方法
    @JsonCreator
    public static User createUser(
        @JsonProperty("id") int id,
        @JsonProperty("name") String name,
        @JsonProperty("email") String email) {
        User user = new User();
        user.id = id;
        user.name = name;
        user.email = email;
        return user;
    }
    
    // getters
    public int getId() { return id; }
    public String getName() { return name; }
    public String getEmail() { return email; }
}
```

## 7. Jackson 配置

### 7.1 ObjectMapper 配置

`ObjectMapper` 类提供了丰富的配置选项，可以自定义其行为：

```java
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import com.fasterxml.jackson.module.paramnames.ParameterNamesModule;

public class ObjectMapperConfigExample {
    public static void main(String[] args) throws Exception {
        ObjectMapper objectMapper = new ObjectMapper();
        
        // 序列化配置
        objectMapper.enable(SerializationFeature.INDENT_OUTPUT); // 格式化输出
        objectMapper.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS); // 日期以字符串形式输出
        objectMapper.disable(SerializationFeature.FAIL_ON_EMPTY_BEANS); // 允许序列化空对象
        
        // 反序列化配置
        objectMapper.enable(DeserializationFeature.ACCEPT_SINGLE_VALUE_AS_ARRAY); // 允许单个值作为数组
        objectMapper.disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES); // 忽略未知属性
        objectMapper.enable(DeserializationFeature.ACCEPT_EMPTY_STRING_AS_NULL_OBJECT); // 空字符串作为 null
        
        // 注册模块
        objectMapper.registerModule(new ParameterNamesModule()); // Java 8 参数名称支持
        objectMapper.registerModule(new JavaTimeModule()); // Java 8 日期时间类型支持
        
        // 测试配置
        User user = new User();
        user.setId(1);
        user.setName("张三");
        
        String json = objectMapper.writeValueAsString(user);
        System.out.println(json);
        
        // 反序列化包含未知字段的 JSON
        String jsonWithUnknown = "{\"id\":1,\"name\":\"张三\",\"unknown_field\":\"value\"}";
        User deserializedUser = objectMapper.readValue(jsonWithUnknown, User.class);
        System.out.println(deserializedUser.getName());
    }
    
    static class User {
        private int id;
        private String name;
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
    }
}
```

### 7.2 全局配置

在 Spring Boot 应用中，可以通过配置类或 application.properties 文件进行 Jackson 的全局配置：

#### 使用 Java 配置类

```java
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.converter.json.Jackson2ObjectMapperBuilder;

@Configuration
public class JacksonConfig {
    
    @Bean
    public ObjectMapper objectMapper(Jackson2ObjectMapperBuilder builder) {
        ObjectMapper objectMapper = builder.build();
        
        // 序列化配置
        objectMapper.enable(SerializationFeature.INDENT_OUTPUT);
        objectMapper.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);
        
        // 反序列化配置
        objectMapper.disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES);
        objectMapper.enable(DeserializationFeature.ACCEPT_SINGLE_VALUE_AS_ARRAY);
        
        // 注册模块
        objectMapper.registerModule(new JavaTimeModule());
        
        return objectMapper;
    }
}
```

#### 使用 application.properties

```properties
# Jackson 配置
spring.jackson.serialization.indent-output=true
spring.jackson.serialization.write-dates-as-timestamps=false
spring.jackson.deserialization.fail-on-unknown-properties=false
spring.jackson.deserialization.accept-single-value-as-array=true
spring.jackson.date-format=yyyy-MM-dd HH:mm:ss
spring.jackson.time-zone=GMT+8
```

## 8. Jackson 高级功能

### 8.1 数据绑定高级特性

#### 8.1.1 泛型类型处理

使用 `TypeReference` 处理泛型类型的反序列化：

```java
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.util.List;
import java.util.Map;

public class GenericTypeExample {
    public static void main(String[] args) throws Exception {
        ObjectMapper objectMapper = new ObjectMapper();
        
        // 反序列化 JSON 数组为 List<User>
        String jsonArray = "[{\"id\":1,\"name\":\"张三\"},{\"id\":2,\"name\":\"李四\"}]";
        List<User> users = objectMapper.readValue(
            jsonArray, 
            new TypeReference<List<User>>() {}
        );
        
        // 反序列化 JSON 对象为 Map<String, User>
        String jsonMap = "{\"user1\":{\"id\":1,\"name\":\"张三\"},\"user2\":{\"id\":2,\"name\":\"李四\"}}";
        Map<String, User> userMap = objectMapper.readValue(
            jsonMap, 
            new TypeReference<Map<String, User>>() {}
        );
        
        // 复杂嵌套泛型
        String complexJson = "{\"data\":[{\"id\":1,\"name\":\"张三\"}],\"total\":1}";
        Response<List<User>> response = objectMapper.readValue(
            complexJson, 
            new TypeReference<Response<List<User>>>() {}
        );
    }
    
    static class User {
        private int id;
        private String name;
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
    }
    
    static class Response<T> {
        private T data;
        private int total;
        
        // getters and setters
        public T getData() { return data; }
        public void setData(T data) { this.data = data; }
        public int getTotal() { return total; }
        public void setTotal(int total) { this.total = total; }
    }
}
```

#### 8.1.2 自定义类型处理器

通过实现 `JsonSerializer` 和 `JsonDeserializer` 接口，或继承 `StdSerializer` 和 `StdDeserializer` 类，创建自定义类型处理器：

```java
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import java.io.IOException;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

// 自定义 LocalDateTime 序列化器
public class LocalDateTimeSerializer extends StdSerializer<LocalDateTime> {
    private static final DateTimeFormatter FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");
    
    public LocalDateTimeSerializer() {
        this(null);
    }
    
    public LocalDateTimeSerializer(Class<LocalDateTime> t) {
        super(t);
    }
    
    @Override
    public void serialize(LocalDateTime value, JsonGenerator gen, SerializerProvider provider) throws IOException {
        gen.writeString(value.format(FORMATTER));
    }
}

// 自定义 LocalDateTime 反序列化器
import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.deser.std.StdDeserializer;
import java.io.IOException;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

public class LocalDateTimeDeserializer extends StdDeserializer<LocalDateTime> {
    private static final DateTimeFormatter FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");
    
    public LocalDateTimeDeserializer() {
        this(null);
    }
    
    public LocalDateTimeDeserializer(Class<?> vc) {
        super(vc);
    }
    
    @Override
    public LocalDateTime deserialize(JsonParser p, DeserializationContext ctxt) throws IOException {
        String value = p.getValueAsString();
        return LocalDateTime.parse(value, FORMATTER);
    }
}

// 注册自定义序列化器和反序列化器
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.module.SimpleModule;
import java.time.LocalDateTime;

public class CustomSerializerExample {
    public static void main(String[] args) throws Exception {
        ObjectMapper objectMapper = new ObjectMapper();
        
        // 创建并配置模块
        SimpleModule module = new SimpleModule();
        module.addSerializer(LocalDateTime.class, new LocalDateTimeSerializer());
        module.addDeserializer(LocalDateTime.class, new LocalDateTimeDeserializer());
        
        // 注册模块
        objectMapper.registerModule(module);
        
        // 测试自定义序列化器
        Event event = new Event();
        event.setTitle("会议");
        event.setStartTime(LocalDateTime.now());
        
        String json = objectMapper.writeValueAsString(event);
        System.out.println(json);
        
        // 测试自定义反序列化器
        Event deserializedEvent = objectMapper.readValue(json, Event.class);
        System.out.println(deserializedEvent.getStartTime());
    }
    
    static class Event {
        private String title;
        private LocalDateTime startTime;
        
        // getters and setters
        public String getTitle() { return title; }
        public void setTitle(String title) { this.title = title; }
        public LocalDateTime getStartTime() { return startTime; }
        public void setStartTime(LocalDateTime startTime) { this.startTime = startTime; }
    }
}
```

### 8.2 树模型高级功能

#### 8.2.1 复杂 JSON 操作

```java
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ArrayNode;
import com.fasterxml.jackson.databind.node.ObjectNode;

public class ComplexTreeExample {
    public static void main(String[] args) throws Exception {
        ObjectMapper objectMapper = new ObjectMapper();
        
        // 创建空的 ObjectNode
        ObjectNode rootNode = objectMapper.createObjectNode();
        
        // 添加基本字段
        rootNode.put("id", 1);
        rootNode.put("name", "项目名称");
        rootNode.put("active", true);
        
        // 添加嵌套对象
        ObjectNode metadataNode = rootNode.putObject("metadata");
        metadataNode.put("createdBy", "admin");
        metadataNode.put("createdAt", "2023-10-01");
        
        // 添加数组
        ArrayNode tagsArray = rootNode.putArray("tags");
        tagsArray.add("重要");
        tagsArray.add("进行中");
        tagsArray.add("高优先级");
        
        // 添加复杂数组元素
        ArrayNode itemsArray = rootNode.putArray("items");
        
        ObjectNode item1 = itemsArray.addObject();
        item1.put("id", 101);
        item1.put("name", "任务 1");
        item1.put("completed", false);
        
        ObjectNode item2 = itemsArray.addObject();
        item2.put("id", 102);
        item2.put("name", "任务 2");
        item2.put("completed", true);
        
        // 转换为格式化的 JSON 字符串
        String formattedJson = objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(rootNode);
        System.out.println(formattedJson);
        
        // 遍历 JSON 树
        System.out.println("\n遍历 JSON 树:");
        traverseJsonNode(rootNode, "");
    }
    
    private static void traverseJsonNode(JsonNode node, String path) {
        if (node.isObject()) {
            node.fieldNames().forEachRemaining(fieldName -> {
                String newPath = path.isEmpty() ? fieldName : path + "." + fieldName;
                traverseJsonNode(node.get(fieldName), newPath);
            });
        } else if (node.isArray()) {
            for (int i = 0; i < node.size(); i++) {
                String newPath = path + "[" + i + "]";
                traverseJsonNode(node.get(i), newPath);
            }
        } else if (node.isValueNode()) {
            System.out.println(path + ": " + node.asText());
        }
    }
}
```

#### 8.2.2 JSON 转换和合并

```java
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ObjectNode;

public class JsonMergeExample {
    public static void main(String[] args) throws Exception {
        ObjectMapper objectMapper = new ObjectMapper();
        
        // 原始 JSON
        String originalJson = "{\"id\":1,\"name\":\"张三\",\"email\":\"zhangsan@example.com\"}";
        JsonNode originalNode = objectMapper.readTree(originalJson);
        
        // 更新数据 JSON
        String updateJson = "{\"name\":\"李四\",\"phone\":\"13800138000\"}";
        JsonNode updateNode = objectMapper.readTree(updateJson);
        
        // 合并 JSON
        JsonNode mergedNode = mergeJson(originalNode, updateNode);
        
        // 输出合并结果
        String mergedJson = objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(mergedNode);
        System.out.println("合并结果:");
        System.out.println(mergedJson);
        
        // 转换 JSON 结构
        JsonNode transformedNode = transformJson(originalNode);
        String transformedJson = objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(transformedNode);
        System.out.println("\n转换结果:");
        System.out.println(transformedJson);
    }
    
    // 合并两个 JSON 节点
    private static JsonNode mergeJson(JsonNode original, JsonNode update) {
        ObjectMapper mapper = new ObjectMapper();
        ObjectNode result = mapper.createObjectNode();
        
        // 添加原始节点的所有字段
        if (original.isObject()) {
            original.fieldNames().forEachRemaining(fieldName -> {
                result.set(fieldName, original.get(fieldName));
            });
        }
        
        // 用更新节点的字段覆盖或添加
        if (update.isObject()) {
            update.fieldNames().forEachRemaining(fieldName -> {
                result.set(fieldName, update.get(fieldName));
            });
        }
        
        return result;
    }
    
    // 转换 JSON 结构
    private static JsonNode transformJson(JsonNode original) {
        ObjectMapper mapper = new ObjectMapper();
        ObjectNode result = mapper.createObjectNode();
        
        // 重新组织字段
        result.put("user_id", original.get("id").asInt());
        result.put("user_name", original.get("name").asText());
        
        // 创建新的嵌套结构
        ObjectNode contactNode = result.putObject("contact");
        contactNode.put("email", original.get("email").asText());
        
        return result;
    }
}
```

### 8.3 性能优化

#### 8.3.1 缓存 ObjectMapper

`ObjectMapper` 实例是线程安全的，应该在应用程序启动时创建并重用：

```java
import com.fasterxml.jackson.databind.ObjectMapper;

public class ObjectMapperSingleton {
    private static final ObjectMapper INSTANCE = new ObjectMapper();
    
    static {
        // 配置 ObjectMapper
        INSTANCE.enable(SerializationFeature.INDENT_OUTPUT);
        INSTANCE.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);
        // 其他配置...
    }
    
    private ObjectMapperSingleton() {}
    
    public static ObjectMapper getInstance() {
        return INSTANCE;
    }
}

// 使用单例 ObjectMapper
public class SingletonUsageExample {
    public static void main(String[] args) throws Exception {
        ObjectMapper objectMapper = ObjectMapperSingleton.getInstance();
        
        User user = new User();
        user.setId(1);
        user.setName("张三");
        
        String json = objectMapper.writeValueAsString(user);
        System.out.println(json);
    }
    
    static class User {
        private int id;
        private String name;
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
    }
}
```

#### 8.3.2 使用 ObjectReader 和 ObjectWriter

对于重复的序列化和反序列化操作，使用 `ObjectReader` 和 `ObjectWriter` 可以提高性能：

```java
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectReader;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ReaderWriterExample {
    public static void main(String[] args) throws Exception {
        ObjectMapper objectMapper = new ObjectMapper();
        
        // 创建可重用的 ObjectWriter
        ObjectWriter writer = objectMapper.writerWithDefaultPrettyPrinter();
        
        // 创建可重用的 ObjectReader
        ObjectReader reader = objectMapper.readerFor(User.class);
        
        // 测试性能
        User user = new User();
        user.setId(1);
        user.setName("张三");
        
        long startTime = System.nanoTime();
        
        // 重复序列化操作
        for (int i = 0; i < 1000; i++) {
            String json = writer.writeValueAsString(user);
        }
        
        long endTime = System.nanoTime();
        System.out.println("序列化 1000 次耗时: " + (endTime - startTime) / 1_000_000 + " ms");
        
        String json = writer.writeValueAsString(user);
        
        startTime = System.nanoTime();
        
        // 重复反序列化操作
        for (int i = 0; i < 1000; i++) {
            User deserializedUser = reader.readValue(json);
        }
        
        endTime = System.nanoTime();
        System.out.println("反序列化 1000 次耗时: " + (endTime - startTime) / 1_000_000 + " ms");
    }
    
    static class User {
        private int id;
        private String name;
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
    }
}
```

#### 8.3.3 流式 API 处理大文件

对于大型 JSON 文件，使用流式 API 可以显著减少内存占用：

```java
import com.fasterxml.jackson.core.JsonFactory;
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.JsonToken;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

public class LargeFileProcessingExample {
    public static void main(String[] args) throws Exception {
        // 生成大型 JSON 文件
        generateLargeJsonFile("large_file.json", 100000);
        
        // 使用流式 API 处理大型 JSON 文件
        processLargeJsonFile("large_file.json");
    }
    
    // 生成大型 JSON 文件
    private static void generateLargeJsonFile(String filePath, int recordCount) throws IOException {
        JsonFactory factory = new JsonFactory();
        
        try (JsonGenerator generator = factory.createGenerator(new FileOutputStream(filePath))) {
            generator.writeStartObject();
            generator.writeFieldName("records");
            generator.writeStartArray();
            
            long startTime = System.currentTimeMillis();
            
            for (int i = 1; i <= recordCount; i++) {
                generator.writeStartObject();
                generator.writeNumberField("id", i);
                generator.writeStringField("name", "用户 " + i);
                generator.writeStringField("email", "user" + i + "@example.com");
                generator.writeNumberField("age", 20 + (i % 50));
                generator.writeEndObject();
                
                // 每 10000 条记录输出进度
                if (i % 10000 == 0) {
                    System.out.println("已生成 " + i + " 条记录");
                }
            }
            
            generator.writeEndArray();
            generator.writeEndObject();
            
            long endTime = System.currentTimeMillis();
            System.out.println("生成文件耗时: " + (endTime - startTime) + " ms");
        }
    }
    
    // 处理大型 JSON 文件
    private static void processLargeJsonFile(String filePath) throws IOException {
        JsonFactory factory = new JsonFactory();
        
        try (JsonParser parser = factory.createParser(new FileInputStream(filePath))) {
            long startTime = System.currentTimeMillis();
            long recordCount = 0;
            long totalAge = 0;
            
            while (!parser.isClosed()) {
                JsonToken token = parser.nextToken();
                
                if (token == null) {
                    break;
                }
                
                if (JsonToken.FIELD_NAME.equals(token) && "records".equals(parser.getCurrentName())) {
                    token = parser.nextToken(); // 进入数组
                    
                    while (!JsonToken.END_ARRAY.equals(token = parser.nextToken())) {
                        if (JsonToken.START_OBJECT.equals(token)) {
                            recordCount++;
                            
                            while (!JsonToken.END_OBJECT.equals(token = parser.nextToken())) {
                                if (JsonToken.FIELD_NAME.equals(token) && "age".equals(parser.getCurrentName())) {
                                    token = parser.nextToken();
                                    totalAge += parser.getValueAsInt();
                                }
                            }
                            
                            // 每 10000 条记录输出进度
                            if (recordCount % 10000 == 0) {
                                System.out.println("已处理 " + recordCount + " 条记录");
                            }
                        }
                    }
                }
            }
            
            long endTime = System.currentTimeMillis();
            
            System.out.println("\n处理结果:");
            System.out.println("总记录数: " + recordCount);
            System.out.println("平均年龄: " + (totalAge / (double) recordCount));
            System.out.println("处理耗时: " + (endTime - startTime) + " ms");
        }
    }
}
```

## 9. Jackson 与 Spring Boot 集成

### 9.1 基本集成

Spring Boot 默认使用 Jackson 作为 JSON 处理库，通过 `spring-boot-starter-web` 依赖自动引入：

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-web</artifactId>
</dependency>
```

### 9.2 RESTful API 示例

```java
import org.springframework.web.bind.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/api/users")
public class UserController {
    private final List<User> users = new ArrayList<>();
    
    public UserController() {
        // 初始化示例数据
        users.add(new User(1, "张三", "zhangsan@example.com"));
        users.add(new User(2, "李四", "lisi@example.com"));
    }
    
    @GetMapping
    public List<User> getAllUsers() {
        return users;
    }
    
    @GetMapping("/{id}")
    public Optional<User> getUserById(@PathVariable int id) {
        return users.stream()
            .filter(user -> user.getId() == id)
            .findFirst();
    }
    
    @PostMapping
    public User createUser(@RequestBody User user) {
        int newId = users.size() > 0 ? users.get(users.size() - 1).getId() + 1 : 1;
        user.setId(newId);
        users.add(user);
        return user;
    }
    
    @PutMapping("/{id}")
    public Optional<User> updateUser(@PathVariable int id, @RequestBody User updatedUser) {
        return users.stream()
            .filter(user -> user.getId() == id)
            .findFirst()
            .map(user -> {
                user.setName(updatedUser.getName());
                user.setEmail(updatedUser.getEmail());
                return user;
            });
    }
    
    @DeleteMapping("/{id}")
    public boolean deleteUser(@PathVariable int id) {
        return users.removeIf(user -> user.getId() == id);
    }
}

class User {
    private int id;
    private String name;
    private String email;
    
    public User() {}
    
    public User(int id, String name, String email) {
        this.id = id;
        this.name = name;
        this.email = email;
    }
    
    // getters and setters
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
}
```

### 9.3 自定义消息转换器

在 Spring Boot 中，可以通过配置自定义的 `MappingJackson2HttpMessageConverter` 来定制 Jackson 的行为：

```java
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.http.converter.json.MappingJackson2HttpMessageConverter;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

import java.util.List;

@Configuration
public class WebConfig implements WebMvcConfigurer {
    
    @Override
    public void configureMessageConverters(List<HttpMessageConverter<?>> converters) {
        // 创建自定义 ObjectMapper
        ObjectMapper objectMapper = new ObjectMapper();
        objectMapper.enable(SerializationFeature.INDENT_OUTPUT);
        objectMapper.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);
        // 其他配置...
        
        // 创建自定义消息转换器
        MappingJackson2HttpMessageConverter converter = new MappingJackson2HttpMessageConverter();
        converter.setObjectMapper(objectMapper);
        
        // 将自定义转换器添加到列表开头，使其优先使用
        converters.add(0, converter);
    }
}
```

### 9.4 全局异常处理

结合 Jackson 实现全局异常处理，返回统一的 JSON 格式响应：

```java
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.context.request.WebRequest;

import java.time.LocalDateTime;

@ControllerAdvice
public class GlobalExceptionHandler {
    
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleGlobalException(Exception ex, WebRequest request) {
        ErrorResponse errorResponse = new ErrorResponse(
            LocalDateTime.now(),
            HttpStatus.INTERNAL_SERVER_ERROR.value(),
            "服务器内部错误",
            ex.getMessage(),
            request.getDescription(false)
        );
        
        return new ResponseEntity<>(errorResponse, HttpStatus.INTERNAL_SERVER_ERROR);
    }
    
    @ExceptionHandler(ResourceNotFoundException.class)
    public ResponseEntity<ErrorResponse> handleResourceNotFoundException(ResourceNotFoundException ex, WebRequest request) {
        ErrorResponse errorResponse = new ErrorResponse(
            LocalDateTime.now(),
            HttpStatus.NOT_FOUND.value(),
            "资源未找到",
            ex.getMessage(),
            request.getDescription(false)
        );
        
        return new ResponseEntity<>(errorResponse, HttpStatus.NOT_FOUND);
    }
    
    @ExceptionHandler(ValidationException.class)
    public ResponseEntity<ErrorResponse> handleValidationException(ValidationException ex, WebRequest request) {
        ErrorResponse errorResponse = new ErrorResponse(
            LocalDateTime.now(),
            HttpStatus.BAD_REQUEST.value(),
            "请求参数验证失败",
            ex.getMessage(),
            request.getDescription(false)
        );
        
        return new ResponseEntity<>(errorResponse, HttpStatus.BAD_REQUEST);
    }
}

class ErrorResponse {
    private LocalDateTime timestamp;
    private int status;
    private String error;
    private String message;
    private String path;
    
    // constructor, getters and setters
    public ErrorResponse(LocalDateTime timestamp, int status, String error, String message, String path) {
        this.timestamp = timestamp;
        this.status = status;
        this.error = error;
        this.message = message;
        this.path = path;
    }
    
    public LocalDateTime getTimestamp() { return timestamp; }
    public void setTimestamp(LocalDateTime timestamp) { this.timestamp = timestamp; }
    public int getStatus() { return status; }
    public void setStatus(int status) { this.status = status; }
    public String getError() { return error; }
    public void setError(String error) { this.error = error; }
    public String getMessage() { return message; }
    public void setMessage(String message) { this.message = message; }
    public String getPath() { return path; }
    public void setPath(String path) { this.path = path; }
}

// 自定义异常类
class ResourceNotFoundException extends RuntimeException {
    public ResourceNotFoundException(String message) {
        super(message);
    }
}

class ValidationException extends RuntimeException {
    public ValidationException(String message) {
        super(message);
    }
}
```

## 10. Jackson 最佳实践

### 10.1 性能最佳实践

1. **重用 ObjectMapper**：ObjectMapper 实例是线程安全的，应该创建一次并在应用程序中重用
2. **使用 ObjectReader 和 ObjectWriter**：对于重复的序列化和反序列化操作，使用这些类可以提高性能
3. **选择适当的 API**：
   - 对于大型 JSON，使用流式 API
   - 对于需要随机访问的 JSON，使用树模型
   - 对于大多数场景，使用数据绑定 API
4. **避免不必要的字段序列化**：使用 `@JsonIgnore` 或 `@JsonInclude` 排除不需要的字段
5. **使用批量操作**：对于大量数据，使用批量处理而不是逐个处理
6. **配置缓存**：启用 Jackson 的缓存功能，如 `MapperFeature.USE_STATIC_TYPING`

### 10.2 安全最佳实践

1. **限制反序列化类**：使用 `ObjectMapper.enableDefaultTyping()` 时要小心，避免反序列化漏洞
2. **验证输入数据**：在反序列化前验证 JSON 数据的格式和内容
3. **使用安全的反序列化设置**：
   ```java
   objectMapper.disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES);
   objectMapper.enable(DeserializationFeature.FAIL_ON_TRAILING_TOKENS);
   ```
4. **避免反序列化敏感类型**：配置 Jackson 只反序列化已知的安全类型
5. **使用最新版本**：保持 Jackson 库的最新版本，以获取安全修复

### 10.3 开发最佳实践

1. **使用有意义的字段名称**：JSON 字段名称应该清晰表达其含义
2. **一致的命名约定**：选择一种命名约定（如驼峰命名法）并在整个应用程序中保持一致
3. **使用数据验证**：结合 Bean Validation API（如 Hibernate Validator）进行数据验证
4. **文档化 API**：使用 Swagger 或 SpringDoc 等工具记录 API 的 JSON 结构
5. **测试序列化和反序列化**：编写单元测试确保 JSON 处理的正确性
6. **处理特殊字符**：确保正确处理包含特殊字符的 JSON 数据
7. **使用版本控制**：当 API 演进时，考虑使用版本控制来管理 JSON 结构的变化

### 10.4 常见问题解决方案

1. **循环引用问题**：
   ```java
   // 使用 @JsonIgnore 打破循环引用
   public class User {
       private int id;
       private String name;
       
       @JsonIgnore
       private Department department;
       
       // getters and setters
   }
   
   // 或使用 @JsonManagedReference 和 @JsonBackReference
   public class User {
       private int id;
       private String name;
       
       @JsonBackReference
       private Department department;
       
       // getters and setters
   }
   
   public class Department {
       private int id;
       private String name;
       
       @JsonManagedReference
       private List<User> users;
       
       // getters and setters
   }
   ```

2. **处理空值和默认值**：
   ```java
   // 使用 @JsonInclude 排除空值
   @JsonInclude(JsonInclude.Include.NON_NULL)
   public class User {
       // fields
   }
   
   // 使用 @JsonSetter 设置默认值
   public class User {
       private int id;
       private String name;
       private int age = 18; // 默认值
       
       @JsonSetter
       public void setAge(Integer age) {
           this.age = (age != null) ? age : 18;
       }
       
       // getters
   }
   ```

3. **处理日期和时间**：
   ```java
   import com.fasterxml.jackson.databind.ObjectMapper;
   import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
   import com.fasterxml.jackson.annotation.JsonFormat;
   import java.time.LocalDateTime;
   
   public class DateExample {
       public static void main(String[] args) throws Exception {
           ObjectMapper objectMapper = new ObjectMapper();
           objectMapper.registerModule(new JavaTimeModule());
           
           Event event = new Event();
           event.setTitle("会议");
           event.setStartTime(LocalDateTime.now());
           
           String json = objectMapper.writeValueAsString(event);
           System.out.println(json);
       }
       
       static class Event {
           private String title;
           
           @JsonFormat(pattern = "yyyy-MM-dd HH:mm:ss")
           private LocalDateTime startTime;
           
           // getters and setters
           public String getTitle() { return title; }
           public void setTitle(String title) { this.title = title; }
           public LocalDateTime getStartTime() { return startTime; }
           public void setStartTime(LocalDateTime startTime) { this.startTime = startTime; }
       }
   }
   ```

4. **自定义序列化和反序列化逻辑**：
   ```java
   import com.fasterxml.jackson.databind.ObjectMapper;
   import com.fasterxml.jackson.databind.module.SimpleModule;
   import java.io.IOException;
   
   public class CustomSerializationExample {
       public static void main(String[] args) throws Exception {
           ObjectMapper objectMapper = new ObjectMapper();
           
           // 创建并注册自定义模块
           SimpleModule module = new SimpleModule();
           module.addSerializer(CustomType.class, new CustomTypeSerializer());
           module.addDeserializer(CustomType.class, new CustomTypeDeserializer());
           objectMapper.registerModule(module);
           
           // 测试自定义序列化和反序列化
           CustomType customType = new CustomType("原始值");
           String json = objectMapper.writeValueAsString(customType);
           System.out.println("序列化结果: " + json);
           
           CustomType deserialized = objectMapper.readValue(json, CustomType.class);
           System.out.println("反序列化结果: " + deserialized.getValue());
       }
       
       static class CustomType {
           private String value;
           
           public CustomType() {}
           
           public CustomType(String value) {
               this.value = value;
           }
           
           public String getValue() { return value; }
           public void setValue(String value) { this.value = value; }
       }
       
       static class CustomTypeSerializer extends com.fasterxml.jackson.databind.JsonSerializer<CustomType> {
           @Override
           public void serialize(CustomType value, com.fasterxml.jackson.core.JsonGenerator gen, com.fasterxml.jackson.databind.SerializerProvider serializers) throws IOException {
               // 自定义序列化逻辑
               gen.writeString("处理后的值: " + value.getValue());
           }
       }
       
       static class CustomTypeDeserializer extends com.fasterxml.jackson.databind.JsonDeserializer<CustomType> {
           @Override
           public CustomType deserialize(com.fasterxml.jackson.core.JsonParser p, com.fasterxml.jackson.databind.DeserializationContext ctxt) throws IOException {
               // 自定义反序列化逻辑
               String value = p.getValueAsString();
               if (value.startsWith("处理后的值: ")) {
                   value = value.substring("处理后的值: ".length());
               }
               return new CustomType(value);
           }
       }
   }
   ```

## 11. Jackson 与其他框架的集成

### 11.1 Spring Boot 集成

Spring Boot 默认使用 Jackson 作为 JSON 处理库，通过 `spring-boot-starter-web` 依赖自动引入。可以通过 `application.properties` 或 `application.yml` 文件进行配置：

```properties
# application.properties
spring.jackson.serialization.indent-output=true
spring.jackson.serialization.write-dates-as-timestamps=false
spring.jackson.deserialization.fail-on-unknown-properties=false
spring.jackson.date-format=yyyy-MM-dd HH:mm:ss
spring.jackson.time-zone=GMT+8
```

```yaml
# application.yml
spring:
  jackson:
    serialization:
      indent-output: true
      write-dates-as-timestamps: false
    deserialization:
      fail-on-unknown-properties: false
    date-format: yyyy-MM-dd HH:mm:ss
    time-zone: GMT+8
```

### 11.2 Hibernate 集成

Jackson 可以与 Hibernate 集成，处理懒加载对象的序列化问题：

```java
import com.fasterxml.jackson.databind.Module;
import com.fasterxml.jackson.datatype.hibernate5.Hibernate5Module;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class HibernateJacksonConfig {
    
    @Bean
    public Module hibernate5Module() {
        Hibernate5Module module = new Hibernate5Module();
        // 配置懒加载对象的处理方式
        module.configure(Hibernate5Module.Feature.FORCE_LAZY_LOADING, false);
        module.configure(Hibernate5Module.Feature.SERIALIZE_IDENTIFIER_FOR_LAZY_NOT_LOADED_OBJECTS, true);
        return module;
    }
}
```

### 11.3 JAX-RS 集成

在 JAX-RS 应用（如 Jersey）中集成 Jackson：

```java
import com.fasterxml.jackson.jaxrs.json.JacksonJsonProvider;
import org.glassfish.jersey.server.ResourceConfig;

public class JaxRsApplication extends ResourceConfig {
    public JaxRsApplication() {
        // 注册 Jackson JSON 提供程序
        register(JacksonJsonProvider.class);
        
        // 注册资源类
        packages("com.example.resources");
    }
}
```

### 11.4 Apache Camel 集成

在 Apache Camel 中使用 Jackson 进行数据格式转换：

```java
import org.apache.camel.builder.RouteBuilder;
import org.apache.camel.model.dataformat.JsonLibrary;

public class CamelJacksonRoute extends RouteBuilder {
    @Override
    public void configure() throws Exception {
        from("direct:marshal")
            .marshal().json(JsonLibrary.Jackson)
            .to("file:output");
            
        from("file:input")
            .unmarshal().json(JsonLibrary.Jackson, User.class)
            .to("direct:process");
    }
    
    static class User {
        private int id;
        private String name;
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
    }
}
```

## 12. Jackson 替代方案

虽然 Jackson 是 Java 生态系统中最流行的 JSON 处理库，但还有其他一些替代方案：

### 12.1 Gson

Gson 是 Google 开发的 JSON 处理库，以其简单的 API 而闻名：

```java
import com.google.gson.Gson;
import java.util.List;

public class GsonExample {
    public static void main(String[] args) {
        Gson gson = new Gson();
        
        // 序列化
        User user = new User(1, "张三", "zhangsan@example.com");
        String json = gson.toJson(user);
        System.out.println(json);
        
        // 反序列化
        User deserializedUser = gson.fromJson(json, User.class);
        System.out.println(deserializedUser.getName());
        
        // 泛型反序列化
        String jsonArray = "[{\"id\":1,\"name\":\"张三\"},{\"id\":2,\"name\":\"李四\"}]";
        List<User> users = gson.fromJson(jsonArray, new com.google.gson.reflect.TypeToken<List<User>>(){}.getType());
        System.out.println("用户数量: " + users.size());
    }
    
    static class User {
        private int id;
        private String name;
        private String email;
        
        public User(int id, String name, String email) {
            this.id = id;
            this.name = name;
            this.email = email;
        }
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
        public String getEmail() { return email; }
        public void setEmail(String email) { this.email = email; }
    }
}
```

### 12.2 JSON-B (JSR 367)

JSON-B 是 Java EE 的 JSON 绑定标准，提供了与 Jackson 类似的功能：

```java
import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;
import java.util.List;
import java.util.ArrayList;

public class JsonBExample {
    public static void main(String[] args) {
        Jsonb jsonb = JsonbBuilder.create();
        
        // 序列化
        User user = new User(1, "张三", "zhangsan@example.com");
        String json = jsonb.toJson(user);
        System.out.println(json);
        
        // 反序列化
        User deserializedUser = jsonb.fromJson(json, User.class);
        System.out.println(deserializedUser.getName());
        
        // 集合序列化和反序列化
        List<User> users = new ArrayList<>();
        users.add(user);
        users.add(new User(2, "李四", "lisi@example.com"));
        
        String jsonArray = jsonb.toJson(users);
        System.out.println(jsonArray);
        
        List<User> deserializedUsers = jsonb.fromJson(jsonArray, new ArrayList<User>(){}.getClass().getGenericSuperclass());
        System.out.println("用户数量: " + deserializedUsers.size());
    }
    
    static class User {
        private int id;
        private String name;
        private String email;
        
        public User() {}
        
        public User(int id, String name, String email) {
            this.id = id;
            this.name = name;
            this.email = email;
        }
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
        public String getEmail() { return email; }
        public void setEmail(String email) { this.email = email; }
    }
}
```

### 12.3 LoganSquare

LoganSquare 是一个专注于性能的 JSON 处理库，使用编译时代码生成来提高性能：

```java
import com.bluelinelabs.logansquare.LoganSquare;

// 需要添加注解
import com.bluelinelabs.logansquare.annotation.JsonField;
import com.bluelinelabs.logansquare.annotation.JsonObject;

@JsonObject
public class LoganSquareExample {
    public static void main(String[] args) throws Exception {
        // 序列化
        User user = new User();
        user.id = 1;
        user.name = "张三";
        user.email = "zhangsan@example.com";
        
        String json = LoganSquare.serialize(user);
        System.out.println(json);
        
        // 反序列化
        User deserializedUser = LoganSquare.parse(json, User.class);
        System.out.println(deserializedUser.name);
    }
    
    @JsonObject
    static class User {
        @JsonField
        int id;
        
        @JsonField
        String name;
        
        @JsonField
        String email;
    }
}
```

### 12.4 Boon

Boon 是一个轻量级的 JSON 处理库，专注于速度和简单性：

```java
import org.boon.json.JsonFactory;
import org.boon.json.ObjectMapper;

public class BoonExample {
    public static void main(String[] args) {
        ObjectMapper mapper = JsonFactory.create();
        
        // 序列化
        User user = new User(1, "张三", "zhangsan@example.com");
        String json = mapper.writeValueAsString(user);
        System.out.println(json);
        
        // 反序列化
        User deserializedUser = mapper.readValue(json, User.class);
        System.out.println(deserializedUser.getName());
    }
    
    static class User {
        private int id;
        private String name;
        private String email;
        
        public User(int id, String name, String email) {
            this.id = id;
            this.name = name;
            this.email = email;
        }
        
        // getters and setters
        public int getId() { return id; }
        public void setId(int id) { this.id = id; }
        public String getName() { return name; }
        public void setName(String name) { this.name = name; }
        public String getEmail() { return email; }
        public void setEmail(String email) { this.email = email; }
    }
}
```

## 13. 结论

Jackson 是一个功能强大、高性能的 Java JSON 处理库，提供了丰富的功能来处理各种 JSON 相关的任务。它的主要优势包括：

- **高性能**：Jackson 是目前性能最高的 JSON 处理库之一，特别适合处理大型 JSON 数据。
- **功能丰富**：提供了数据绑定、树模型、流式 API 等多种处理 JSON 的方式，满足不同场景的需求。
- **灵活性**：通过丰富的注解和自定义选项，可以灵活地控制序列化和反序列化行为。
- **广泛的生态系统**：支持多种数据格式（JSON、XML、YAML、CSV 等）和各种 Java 技术栈。
- **良好的社区支持**：作为 Java 生态系统中最流行的 JSON 处理库之一，拥有活跃的社区和丰富的资源。

Jackson 与 Spring Boot 等主流框架的无缝集成，使其成为 Java 开发者处理 JSON 数据的首选工具。无论是构建 RESTful API、处理配置文件，还是进行数据交换，Jackson 都提供了简单而强大的解决方案。

对于 Java 开发者来说，掌握 Jackson 的使用技巧，不仅可以提高开发效率，还可以优化应用程序的性能和用户体验。通过合理使用 Jackson 的各种功能和最佳实践，可以构建出更加高效、可靠的 Java 应用程序。

## 14. 参考资料

- [Jackson 官方文档](https://github.com/FasterXML/jackson-docs)
- [Jackson GitHub 仓库](https://github.com/FasterXML/jackson)
- [Baeldung Jackson 教程](https://www.baeldung.com/jackson)
- [Spring Boot 官方文档](https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/)
- [JSON 官方网站](https://www.json.org/)
