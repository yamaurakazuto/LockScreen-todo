# Lock Screen Todo アプリケーション仕様書

## 1. プロジェクト概要
iPhoneのロック画面から直接操作可能なTodoリストアプリケーションを開発します。

### 1.1 使用技術
- フロントエンド: React Native
- バックエンド: Go
- ロック画面統合: iOS Widgets (WidgetKit)

## 2. システムアーキテクチャ

### 2.1 全体構成
```
[iOS App (React Native)]
    ↕ REST API
[Go Backend Server]
    ↕
[Database]

[iOS Widget (Swift)]
    ↕ App Groups (共有データ)
[iOS App (React Native)]
```

### 2.2 主要コンポーネント
- React Nativeアプリ（メインアプリ）
- Goバックエンド（API・データ管理）
- iOSウィジェット（ロック画面表示）
- ローカルストレージ（オフライン対応）

## 3. 機能要件

### 3.1 コア機能

#### 3.1.1 Todo管理（アプリ内）
- Todoの作成
- Todoの編集
- Todoの削除
- Todo完了/未完了の切り替え
- Todoの並び替え

#### 3.1.2 ロック画面ウィジェット
- 最新のTodo（最大3-5件）をロック画面に表示
- タップでTodoの完了/未完了を切り替え
- アプリへのディープリンク

#### 3.1.3 データ同期
- バックエンドとの自動同期
- オフライン時のローカル保存
- オンライン復帰時の同期

## 4. 技術仕様

### 4.1 React Nativeアプリ

#### 4.1.1 画面構成
- ホーム画面
  - Todoリスト表示
  - 追加ボタン
  - 並び替え機能
- Todo作成/編集画面
  - タイトル入力
  - 詳細入力（オプション）
  - 優先度設定（オプション）
- 設定画面
  - ウィジェット設定
  - 同期設定

#### 4.1.2 必要なライブラリ
- `@react-navigation/native` (画面遷移)
- `@react-native-async-storage/async-storage` (ローカルストレージ)
- `axios` (API通信)
- `react-native-reanimated` (アニメーション)

#### 4.1.3 ネイティブブリッジ（Swift連携）
- ウィジェット用のデータ共有
- App Groupsを使用したデータ共有
- UserDefaultsでの軽量データ保存

### 4.2 Goバックエンド

#### 4.2.1 API仕様
ベースURL: `https://api.example.com/v1`

| メソッド | エンドポイント | 概要 |
| --- | --- | --- |
| GET | `/todos` | 全Todo取得 |
| POST | `/todos` | Todo作成 |
| PUT | `/todos/:id` | Todo更新 |
| DELETE | `/todos/:id` | Todo削除 |
| PATCH | `/todos/:id/toggle` | 完了/未完了切り替え |

**Todo作成リクエストボディ**
```json
{ "title": "string", "description": "string", "priority": 1 }
```

**Todo更新リクエストボディ**
```json
{ "title": "string", "description": "string", "priority": 1, "completed": false }
```

#### 4.2.2 データモデル
```go
type Todo struct {
    ID          string    `json:"id"`
    UserID      string    `json:"user_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Priority    int       `json:"priority"` // 1-3
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

#### 4.2.3 使用フレームワーク・ライブラリ
- `gin-gonic/gin` (Webフレームワーク)
- `gorm.io/gorm` (ORM)
- PostgreSQL or SQLite (データベース)
- `golang-jwt/jwt` (認証)

### 4.3 iOSウィジェット（Swift/WidgetKit）

#### 4.3.1 ウィジェット仕様
- **サイズ**: Small, Medium
- **更新頻度**: 15分ごと（iOS制限）
- **表示内容**:
  - Small: 最新Todo 1件
  - Medium: 最新Todo 3-5件

#### 4.3.2 データ共有
- App Groupsを使用
- UserDefaults共有でTodoデータを保存
- JSON形式でシリアライズ

## 5. データフロー

### 5.1 Todo作成フロー
```
1. ユーザーがアプリでTodo作成
2. ローカルDBに即座に保存
3. バックエンドAPIにPOSTリクエスト
4. 成功時: サーバーIDでローカル更新
5. ウィジェット用データを共有ストレージに書き込み
6. ウィジェットが次回更新時に反映
```

### 5.2 ロック画面からの操作フロー
```
1. ユーザーがウィジェットをタップ
2. ウィジェットがディープリンクでアプリ起動
3. アプリが該当Todoを完了状態に変更
4. バックエンドに同期
5. 共有ストレージ更新
```

## 6. セキュリティ

### 6.1 認証
- JWT（JSON Web Token）ベースの認証
- トークンの有効期限: 7日間
- リフレッシュトークン機能

### 6.2 通信
- HTTPS必須
- API Rate Limiting実装

## 7. 開発フェーズ

### Phase 1: MVP（最小機能）
- React NativeでのTodo CRUD機能
- Go APIの基本実装
- ローカルストレージ

### Phase 2: ウィジェット統合
- iOSウィジェット開発
- App Groups連携
- データ共有機能

### Phase 3: 同期・最適化
- バックエンド同期
- オフライン対応
- パフォーマンス最適化

## 8. 制約事項

### 8.1 iOS制約
- ウィジェットの更新頻度制限（15分）
- ウィジェットからの直接API通信不可
- App Groups経由のデータ共有が必須

### 8.2 技術的制約
- React Nativeからネイティブコード（Swift）への連携が必要
- ウィジェット部分は純粋なSwiftで実装
