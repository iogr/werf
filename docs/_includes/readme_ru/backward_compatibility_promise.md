<!-- THIS FILE IS AUTOGENERATED BY werf docs COMMAND! DO NOT EDIT! -->

> _Note:_ Настоящее обещание относится к werf, начиная с версии 1.0, и не относится к предыдущим версиям или версиям dapp

werf использует [семантическое версионирование](https://semver.org/lang/ru/). Это значит, что мажорные версии (1.0, 2.0) могут быть обратно не совместимыми между собой. В случае werf это означает, что обновление на следующую мажорную версию _может_ потребовать полного передеплоя приложений, либо других ручных операций.

Минорные версии (1.1, 1.2, etc) могут добавлять новые "значительные" изменения, но без существенных проблем обратной совместимости в пределах мажорной версии. В случае werf это означает, что обновление на следующую минорную версию в большинстве случаев будет беспроблемным, но _может_ потребоваться запуск предоставленных скриптов миграции.

Патч-версии (1.1.0, 1.1.1, 1.1.2) могут добавлять новые возможности, но без каких-либо проблем обратной совместимости в пределах минорной версии (1.1.x).
В случае werf это означает, что обновление на следующий патч (следующую патч-версию) не должно вызывать проблем и требовать каких-либо ручных действий.

- Мы **не гарантируем** обратную совместимость между версиями:
  - канала обновлений `alpha`,
  - канала обновлений `beta`,
  - канала обновлений `ea`.
- Мы **гарантируем** обратную совместимость между версиями:
  - канала обновлений `stable` в пределах минорной версии (1.1.x),
  - канала обновлений `rock-solid` в пределах минорной версии (1.1.x).
