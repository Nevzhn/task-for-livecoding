# task-for-livecoding
В complete решенные задачи
в tasks нерешенные 

task 1 - https://easyoffer.ru/live-code-tasks/golang-developer/poisk-oshibki-v-kode-1
Необходимо внести исправления в код, который генерирует случайные числа от 0 до 9, 
выбирает среди них повторяющиеся значения, проверяет их на уникальность и отправляет уникальные числа в канал

task 2 - https://easyoffer.ru/live-code-tasks/golang-developer/kastomnyie-oshibki-1
Проанализировать поведение при сравнении возвращаемых значений с nil и определить, что будет выведено в консоль

task 3 - https://cloud.mail.ru/public/VYnE/rcu1i4k1j/%5BSupersliv.biz%5D%206.%20Concurrency/6.1.%20%D0%97%D0%B0%D0%B4%D0%B0%D1%87%D0%B0%20%D0%BD%D0%B0%20%D0%BD%D0%B0%D0%BF%D0%B8%D1%81%D0%B0%D0%BD%D0%B8%D0%B5%20%D0%BE%D0%B1%D0%B5%D1%80%D1%82%D0%BA%D0%B8%20%D0%BD%D0%B0%D0%B4%20%D0%B4%D0%BE%D0%BB%D0%B3%D0%B8%D0%BC%20%D1%81%D0%B5%D1%82%D0%B5%D0%B2%D1%8B%D0%BC%20%D0%B2%D1%8B%D0%B7%D0%BE%D0%B2%D0%BE%D0%BC.mp4
Требуется доработать метод SimulateRequest так, чтобы:
- Код работал в конкурентной среде
- При долгом ожидании запрос отваливался по таймауту
- На консоль печаталось время выполнения запроса