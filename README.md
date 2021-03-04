# The Black Pearl: 黑珍珠号



![](https://s2.ax1x.com/2019/10/27/KyNZTK.png)

**blackpearl** is a personal helper written in golang, it automatic many things for me in daily life. Currently it do those things:

- [x] Support hexo blog templates;
- [x] Automatically upload images for me. (this powers by picbed, upload 2 different picbed);
- [ ] Connect with my work cloud (db.tsuai.cn), I can check my todos here (which synchronized with my Phone), using https://github.com/gizak/termui to do this job;
- [ ] GUI support for customizable GUI in terminal.
- [ ] Message tunnel in Colibri, waiting for bring this feature up.



## Updates

- **2050.01.01**: ..

- **2021.02.03**: I decided to make a terminal GUI for blackpearl, at blackpearl, I can monitor a lot of things, such as my todos, stocks, weather and etc. These components can be turned on or off and highly customizable. Especially for **todo**, we not only need show what to do next, but also need add/edit/delete todos, this function can be connected with daybreak;

  



## Build

If you have golang enviroment, you can build binary yourself:

```
./build_blackpearl.sh
```




## Usage

*blackpearl* is module design, it current support below modules:

- `blog`: helps me generates blog templates:

  ```
  blackpearl blog `some title`
  ```

- `upload`: helps me upload images or files:

  ```
  blackpearl upload -p /image/path.jpg
  ```

  todo in this module:

  - [ ] Support directly upload image on clipboard;

- `push`: this will push a text content or an image to uranus message system, you can receive a message on Chats app which can be downloaded from: http://loliloli.pro

- `code`: generates C++ projects;

- `scrap`: scrap website (to be done);

- `git`: quickly do some git commands;

- more to be add, PR is welcome.



For more detail, type:

```
blackpearl -h
```



## Todos

I want built **blackpearl** into a powerful golang utils which means we should plan some functions for it:

- [ ] do something like scrap;
- [x] update picbed package
- [x] integrate with my ultimate uranus universe;
- [ ] Add `blackpearl daybreak login` to login on daybreak;
- [ ] Add todo panel, and interactive on todo status, as well as editing on every single todos.

## Copyright

All rights belongs to Fagang Jin, code released under Apache Lincense.


