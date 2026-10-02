# frontend-lab

本仓库是前端学习项目的**聚合仓**：用 **Git 子模块** 引用 [`wxz6289`](https://github.com/wxz6289) 下各独立仓库，父仓只保存**子模块 commit 指针**，日常维护在各自子仓库中进行。

---

## 2. 子模块列表（按主题）

### 2.1 基础与语言

| 本仓目录 | 远程仓库 | 说明 |
|----------|----------|------|
| `html-css/` | [wxz6289/html-css](https://github.com/wxz6289/html-css) | HTML / CSS |
| `js/` | [wxz6289/js](https://github.com/wxz6289/js) | JavaScript |
| `typescript/` | [wxz6289/learn-typescript](https://github.com/wxz6289/TypeScript) | TypeScript |

### 2.2 框架与 SSR

| 本仓目录 | 远程仓库 | 说明 |
|----------|----------|------|
| `react/` | [wxz6289/react](https://github.com/wxz6289/react) | React |
| `vue/` | [wxz6289/vue](https://github.com/wxz6289/vue) | Vue |
| `next/` | [wxz6289/next](https://github.com/wxz6289/next) | Next.js |
| `nuxt/` | [wxz6289/nuxtjs](https://github.com/wxz6289/nuxtjs) | Nuxt |

### 2.3 构建工具

| 本仓目录 | 远程仓库 | 说明 |
|----------|----------|------|
| `webpack/` | [wxz6289/webpack](https://github.com/wxz6289/webpack) | Webpack |
| `vite/` | [wxz6289/vite](https://github.com/wxz6289/vite) | Vite |

### 2.4 桌面端

| 本仓目录 | 远程仓库 | 说明 |
|----------|----------|------|
| `electron/` | [wxz6289/electron](https://github.com/wxz6289/electron) | Electron |
| `tauri/` | [wxz6289/tauri](https://github.com/wxz6289/tauri) | Tauri |

### 2.5 图形与响应式

| 本仓目录 | 远程仓库 | 说明 |
|----------|----------|------|
| `rxjs/` | [wxz6289/rxjs](https://github.com/wxz6289/rxjs) | RxJS |

> **Three.js**：原 `threejs` 子模块已从本聚合仓移除（仓库体积大）。需要时单独克隆 [wxz6289/threejs](https://github.com/wxz6289/threejs)。

### 2.6 跨端与其它

| 本仓目录 | 远程仓库 | 说明 |
|----------|----------|------|
| `uniapp/` | [wxz6289/uni-app](https://github.com/wxz6289/uni-app) | uni-app |
| `flutter/` | [wxz6289/flutter](https://github.com/wxz6289/flutter) | Flutter（本地 `learn` 下常为目录名 `fultter`） |
| `admin/` | [wxz6289/admin](https://github.com/wxz6289/admin) | 管理后台示例 |
| `frontend-interview/` | [wxz6289/frontend-interview](https://github.com/wxz6289/frontend-interview) | 前端面试资料 |

### 2.7 微前端

| 本仓目录 | 远程仓库 | 说明 |
|----------|----------|------|
| `qiankun/` | [wxz6289/qiankun](https://github.com/wxz6289/qiankun) | qiankun 2.x 主应用 + React / Vue 子应用 |

**本地运行微前端示例：**

```bash
cd qiankun
npm install
npm run dev
```

浏览器打开 <http://localhost:7100>（主应用 `7100`，React `7101`，Vue `7102`）。详见子仓 [qiankun/README.md](qiankun/README.md)。

---

## 3. 克隆与子模块检出

**一次性拉取父仓及全部子模块：**

```bash
git clone --recurse-submodules git@github.com:wxz6289/frontend-lab.git
cd frontend-lab
```

**已克隆父仓、子模块目录仍为空时：**

```bash
cd frontend-lab
git submodule update --init --recursive
```

**只检出部分子模块（例如 html-css + qiankun）：**

```bash
git submodule update --init html-css qiankun
```

---

## 4. 在子模块里改代码并回写父仓

1. 进入某个子模块目录，在**该子模块默认分支**上开发（各仓可能是 `main` 或 `master`，以子仓为准）。
2. 在子仓内 `commit` 并 `push` 到子仓远程。
3. 回到 `frontend-lab` 根目录，`git add <子模块目录>` 更新指针，再 `commit` 并 `push` 父仓。

示例（以 `react` 为例，分支名请按子仓实际修改）：

```bash
cd react
git checkout main
git add -A && git commit -m "feat: ..." && git push origin HEAD

cd ..
git add react
git commit -m "chore: bump react submodule"
git push origin main
```

---

## 5. 相关链接

- 聚合仓：[wxz6289/frontend-lab](https://github.com/wxz6289/frontend-lab)
