# SVG 动画效果

- 在 SVG 中，实现动画效果主要依赖于 SMIL (Synchronized Multimedia Integration Language) 动画规范。
- `<animate>`、`<animateTransform>` 和 `<animateMotion>` 是 SMIL 中最核心的三个动画标签。

## `<animate>` 基础属性动画

- **作用：** 用于对 SVG 元素的任意属性值（如颜色、透明度、半径、坐标等）进行随时间变化的动画。
- **核心属性：**
  - `attributeName`：要发生变化的属性名（如 r, fill, opacity, cx）。
  - `from / to`：属性的起始值和结束值。
  - `values`：用分号 ; 隔开的多个值，定义动画的完整关键帧（如果使用了 values，则忽略 from 和 to）。

- **示例：** 呼吸灯效果（改变半径和透明度）

```xml
<svg xmlns="http://www.w3.org/2000/svg" width="200" height="200" viewBox="0 0 200 200">
  <circle cx="100" cy="100" r="20" fill="#ff5722">
    <!-- 半径从 20 变到 50 再变回 20 -->
    <animate attributeName="r" values="20; 50; 20" dur="2s" repeatCount="indefinite" />

    <!-- 透明度从 1 变到 0.3 再变回 1 -->
    <animate attributeName="opacity" values="1; 0.3; 1" dur="2s" repeatCount="indefinite" />
  </circle>
</svg>
```

## `<animateTransform>` 变换动画

- **作用：** 专门用于对元素的 transform（变换）属性进行动画处理。
   - 支持平移（translate）、旋转（rotate）、缩放（scale）和倾斜（skewX/skewY）。
- **核心属性：**
  - `attributeName`：固定为 "transform"。
  - `type`：变换类型（translate, rotate, scale, skewX, skewY）。
  - `from / to / values`：注意，这里的值格式与 CSS 中的 transform 略有不同，通常需要包含坐标参数。

- **示例：** 自转的矩形（旋转）与 弹跳效果（平移）

```xml
<svg xmlns="http://www.w3.org/2000/svg" width="300" height="200" viewBox="0 0 300 200">
  <!-- 1. 旋转动画 -->
  <rect x="40" y="40" width="40" height="40" fill="#2196f3">
    <!-- 注意：rotate 的 from/to 需要指定旋转中心点 (角度 x y) -->
    <!-- 这里以矩形中心 (60, 60) 为旋转中心 -->
    <animateTransform attributeName="transform" 
                      type="rotate" 
                      from="0 60 60" 
                      to="360 60 60" 
                      dur="3s" 
                      repeatCount="indefinite" />
  </rect>

  <!-- 2. 平移（弹跳）动画 -->
  <circle cx="200" cy="100" r="15" fill="#4caf50">
    <!-- 在 Y 轴上上下移动 -->
    <animateTransform attributeName="transform" 
                      type="translate" 
                      values="0,0; 0,-40; 0,0" 
                      dur="1s" 
                      repeatCount="indefinite" />
  </circle>
</svg>
```

## `<animateMotion>` 路径运动动画

- **作用：** 让元素沿着一条指定的路径（Path） 进行运动。这是 SVG 动画中最强大、最独特的功能之一。
- **核心属性：**
  - `path`：定义运动轨迹的 SVG 路径数据（与 <path> 的 d 属性语法一致）。
  - `rotate`：元素在运动时是否跟随路径的切线方向旋转（auto 跟随，auto-reverse 反向跟随，0 或不写则保持原方向）。
  - `<mpath>`：子标签，用于引用外部已经定义好的 <path> 元素，避免重复写路径代码。

- **示例：** 沿曲线飞行的小飞机

```xml
<svg xmlns="http://www.w3.org/2000/svg" width="400" height="200" viewBox="0 0 400 200">
  <!-- 定义一条不可见的运动轨迹 -->
  <path id="flightPath" d="M 20 150 C 100 20, 300 180, 380 50" stroke="#ccc" stroke-dasharray="5,5" fill="none" />
  
  <!-- 运动的小球（代表飞机） -->
  <circle r="8" fill="#e91e63">
    <!-- 方式一：直接写 path -->
    <!-- <animateMotion dur="4s" repeatCount="indefinite" path="M 20 150 C 100 20, 300 180, 380 50" rotate="auto" /> -->
    
    <!-- 方式二：使用 mpath 引用外部路径（推荐） -->
    <animateMotion dur="4s" repeatCount="indefinite" rotate="auto">
      <mpath href="#flightPath" />
    </animateMotion>
  </circle>
</svg>
```

## 公共控制属性

以上三个标签都共享一组控制动画时间、节奏和状态的核心属性：
- `dur (Duration)`：动画单次播放的持续时间。如 2s, 500ms。
- `repeatCount`：重复次数。可以是具体数字（如 3），或 **indefinite**（无限循环）。
- `begin`：动画开始时间。
  - `时间`：0s（默认）, 2s（延迟2秒）。
  - `事件`：click（点击时触发）, mouseover。
  - `其他动画`：anim1.end（在 id 为 anim1 的动画结束时触发，用于做动画链）。
- `fill`：动画结束后的状态。
  - `remove（默认）`：动画结束后恢复到初始状态。
  - `freeze`：动画结束后保持在最后一帧的状态。
- `calcMode & keySplines`：控制动画的缓动效果（类似 CSS 的 transition-timing-function）。
  - calcMode="spline" 时，配合 keySplines="0.42 0 0.58 1" 可以实现类似 CSS ease-in-out 的效果。

## 组合使用

在实际开发中，我们经常将这几个标签嵌套在同一个图形中，实现复杂的复合动画。

- **示例：** 一个变色、旋转且沿路径运动的星星

```xml
<svg xmlns="http://www.w3.org/2000/svg" width="400" height="200" viewBox="0 0 400 200">
  <path id="starPath" d="M 10 100 Q 200 10 390 100" stroke="#eee" fill="none" />
  
  <g>
    <!-- 1. 沿路径运动 -->
    <animateMotion dur="5s" repeatCount="indefinite">
      <mpath href="#starPath" />
    </animateMotion>
    
    <!-- 星星图形 -->
    <polygon points="12,2 15,9 22,9 16,14 18,21 12,17 6,21 8,14 2,9 9,9" fill="gold">
      <!-- 2. 自转 -->
      <animateTransform attributeName="transform" type="rotate" from="0 12 12" to="360 12 12" dur="2s" repeatCount="indefinite" />
      <!-- 3. 颜色渐变 -->
      <animate attributeName="fill" values="gold;orange;red;gold" dur="3s" repeatCount="indefinite" />
    </polygon>
  </g>
</svg>
```
