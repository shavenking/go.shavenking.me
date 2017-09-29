# Semantic Versioning 2.0.0 (SemVer)

注意：以下內容都是我個人看法，原文請看 [semver.org](http://semver.org)。

## 總結

`MAJOR.MINOR.PATCH`

1. `MAJOR` 當新的功能無法向下相容時。
2. `MINOR` 當新的功能可以向下相容時。
3. `PATCH` 修正現有的 BUG，且該修正能向下相容。

這邊指得是 API 還是一般系統也適用？API 在系統來說可能是公開的功能（使用者接觸的），在程式套件來說就是指套件提供的功能，在 Restful API 就是指 API？哈哈哈我在說三Ｘ。

預先建制、釋出的版本，可以以擴充的方式加入這個版本命名機制，比方說 alpha beta？

## 介紹

Dependency Hell

隨著系統成長，相依套件會越來越多。

當一個系統有複雜的相依性時，推出新的版本會造成極大的困擾。

Dependency Hell 就是指因為相依性太高，在升級系統版本時，會受限於相依套件的版本，包含因為版本限制而無法升級，或者是因為版本混亂導致程式錯誤。

因此作者提出了這樣的版本命名方式，前提是系統必需先定義出 API。

一但確定 API 之後，就可以依循這個命名規則。

## 規格（規定？）

### 1. 必須定義明確的 API

> Software using Semantic Versioning MUST declare a public API. This API could be declared in the code itself or exist strictly in documentation. However it is done, it should be precise and comprehensive.

### 2. `X.Y.Z` 三個數字的規定

- 正整數
- 不可有 0 開頭
- 必須遞增（感覺是 +1 +1，不能一次跳 +2）？

 > A normal version number MUST take the form X.Y.Z where X, Y, and Z are non-negative integers, and MUST NOT contain leading zeroes. X is the major version, Y is the minor version, and Z is the patch version. Each element MUST increase numerically. For instance: 1.9.0 -> 1.10.0 -> 1.11.0.

### 3. 一但版本釋出後，任何修改都必須遞增版號

> Once a versioned package has been released, the contents of that version MUST NOT be modified. Any modifications MUST be released as a new version.

### 4. `0.Y.Z` 代表 API 還在開發中，視為不穩定的版本

> Major version zero (0.y.z) is for initial development. Anything may change at any time. The public API should not be considered stable.

### 5. `1.0.0` 視為正式版本，接下來須依循此命名規則？

> Version 1.0.0 defines the public API. The way in which the version number is incremented after this release is dependent on this public API and how it changes.

### 6. Patch 版號只有在可以向下相容的 BUG 修復時增加，BUG 是指修正不正確的程式行為

> Patch version Z (x.y.Z | x > 0) MUST be incremented if only backwards compatible bug fixes are introduced. A bug fix is defined as an internal change that fixes incorrect behavior.

### 7. Minor 版號遞增時機、規則

- 新增任何可以向下相容的功能時
- 任何現有的功能被標記為 Deprecated 時
- （非必須）內部程式有顯著的功能、改善時
- （非必須）可以包涵 Patch 層級的修改
- 當 Minor 遞增時，Patch 需歸零

> Minor version Y (x.Y.z | x > 0) MUST be incremented if new, backwards compatible functionality is introduced to the public API. It MUST be incremented if any public API functionality is marked as deprecated. It MAY be incremented if substantial new functionality or improvements are introduced within the private code. It MAY include patch level changes. Patch version MUST be reset to 0 when minor version is incremented.

### 8. Major 版號遞增時機、規則

- 新增任何無法向下相容的功能、修正時
- （非必須）可以包涵 Minor、Patch 層級的修改
- 當 Major 遞增時，Minor、Patch 需歸零

> Major version X (X.y.z | X > 0) MUST be incremented if any backwards incompatible changes are introduced to the public API. It MAY include minor and patch level changes. Patch and minor version MUST be reset to 0 when major version is incremented.

### 9. 預先釋出的版本號規則

- 預先釋出的版本號可以加在 Patch 後面，只能包含 `[0-9A-Za-z-]` 符號
- 預先釋出的版本號不得為空、不得為 0 開頭
- 預先釋出的版本號代表目前的 API 不穩定

> A pre-release version MAY be denoted by appending a hyphen and a series of dot separated identifiers immediately following the patch version. Identifiers MUST comprise only ASCII alphanumerics and hyphen [0-9A-Za-z-]. Identifiers MUST NOT be empty. Numeric identifiers MUST NOT include leading zeroes. Pre-release versions have a lower precedence than the associated normal version. A pre-release version indicates that the version is unstable and might not satisfy the intended compatibility requirements as denoted by its associated normal version. Examples: 1.0.0-alpha, 1.0.0-alpha.1, 1.0.0-0.3.7, 1.0.0-x.7.z.92.

### 10. 建置的版本號規則

- 需串在 patch 或 pre-release 版本號之後，需以 + 號開頭，只能包含 `[0-9A-Za-z-]` 符號
- 不得為空
- 決定版本優先順序時，通常會被忽略

> Build metadata MAY be denoted by appending a plus sign and a series of dot separated identifiers immediately following the patch or pre-release version. Identifiers MUST comprise only ASCII alphanumerics and hyphen [0-9A-Za-z-]. Identifiers MUST NOT be empty. Build metadata SHOULD be ignored when determining version precedence. Thus two versions that differ only in the build metadata, have the same precedence. Examples: 1.0.0-alpha+001, 1.0.0+20130313144700, 1.0.0-beta+exp.sha.5114f85.

### 11. 版本號優先順序規則

- 優先順序指得是比較不同版號時的排序
- 計算優先順序時，必須將 Major、Minor、Patch、Pre-release 分開計算（建置版號通常不列入考量），由左到右，以數字大小來比較
- 範例：1.0.0 < 2.0.0 < 2.1.0 < 2.1.1
- Pre-release 版號優先順序較小，範例：1.0.0-alpha < 1.0.0
- Pre-release 版號比較方式為 ASCII 排序，範例：1.0.0-alpha < 1.0.0-alpha.1 < 1.0.0-alpha.beta < 1.0.0-beta < 1.0.0-beta.2 < 1.0.0-beta.11 < 1.0.0-rc.1 < 1.0.0

> Precedence refers to how versions are compared to each other when ordered. Precedence MUST be calculated by separating the version into major, minor, patch and pre-release identifiers in that order (Build metadata does not figure into precedence). Precedence is determined by the first difference when comparing each of these identifiers from left to right as follows: Major, minor, and patch versions are always compared numerically. Example: 1.0.0 < 2.0.0 < 2.1.0 < 2.1.1. When major, minor, and patch are equal, a pre-release version has lower precedence than a normal version. Example: 1.0.0-alpha < 1.0.0. Precedence for two pre-release versions with the same major, minor, and patch version MUST be determined by comparing each dot separated identifier from left to right until a difference is found as follows: identifiers consisting of only digits are compared numerically and identifiers with letters or hyphens are compared lexically in ASCII sort order. Numeric identifiers always have lower precedence than non-numeric identifiers. A larger set of pre-release fields has a higher precedence than a smaller set, if all of the preceding identifiers are equal. Example: 1.0.0-alpha < 1.0.0-alpha.1 < 1.0.0-alpha.beta < 1.0.0-beta < 1.0.0-beta.2 < 1.0.0-beta.11 < 1.0.0-rc.1 < 1.0.0.
  
## 為什麼需要這樣的版號命名規則？

也許大家都已經有類似的版號命名規則，但是有一份完整的規範，會讓整個軟體環境更好。

有了這樣的規範，在使用相依套件時，就可以很安心的依照版號使用。

## FAQ

### 0.y.z 的命名規則？

以 0.1.0 當作起始命名，並依序遞增 Minor

### 何時為 1.0.0 釋出時機？

- 已經在正式環境上跑的系統
- 已經有人在使用的套件
- 已經開始擔心會不會有向下相容問題的程式

### 這樣的版號規則會降低開發效率嗎？

一般來說 API 改變的幅度不會太大，如果目前系統每天都還有大幅改變的話，應該是還在 0.Y.Z 的階段。

### 如果每次有無法向下相容的功能推出，都必須增加 Major 版號，那 Major 版號會不會很快就破百？

正常來說一個穩定的系統，不應該有大量無法向下相容的功能推出。如果有，就應該去反省為什麼會有，哈哈哈哈

### 如果不小心釋出了一個無法向下相容的 Minor 版本，該怎麼辦？

注意，即使是這樣的狀況，也不能在不遞增版號的狀況下，修改程式，取而代之的應該是修正程式，並遞增新的 Minor 版號，讓整個循環恢復，並通知使用者版號的問題。

### 如果只是更新我的套件的相依套件版本，我需要遞增版號嗎？

也許使用我的套件的軟體，也有相依在我相依的套件上，因此，看我更新套件是為了修復還是新增功能，來決定要遞增哪個層級的版號，結論就是還是要遞增版號！

### 如果不小心改了 API 且沒有依照這個規則遞增版號？

八七喔就說要照這個版號世界才會美好，還要幹這種事？

### 怎麼處理準備淘汰的功能？

遞增 Minor 版號並釋出準備淘汰的訊息，並在下一版 Major 版號移除，注意，在 Major 版號釋出之前必須有至少一個 Minor 版號釋出準備淘汰的訊息。

### 版號有大小限制嗎？

沒有，但是 255 已經夠用了。
