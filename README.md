# 気分記録アプリケーション「lantana」
lantanaは気分記録アプリケーションです。  
![lantana](https://raw.githubusercontent.com/mt3hr/lantana/main/document/img/lantana.png)  

## ダウンロード
[lantanaダウンロード](https://github.com/mt3hr/lantana/releases/latest)  

## 実行
「lantana.exe」または「lantana_server.exe」をダブルクリック  
（lantana.server.exeの場合は起動後「[http://localhost:7777](http://localhost:7777)」にアクセス）  

<details>
<summary>開発者向け</summary>

開発者向けと言いつつ自分向けです。  
ビルドに必要パッケージを公開していないのでビルド不可能だと思います。  

### 開発環境

### セットアップ
1. Golang バージョン1.20の開発環境を用意する  
2. Cコンパイラを用意する（cgo使用のため）  
3. Node.js バージョン18.12.1の開発環境を用意する  
4. 以下のコマンドを実行する  
```
npm i
```

### ビルド・インストール

アプリケーションインストール  
```
npm run go_mod
npm run install_app
```

サーバインストール  
```
npm run go_mod
npm run install_build
```
</details>