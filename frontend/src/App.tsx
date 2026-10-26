import { useState, useEffect } from 'react'
// Убедись, что путь правильный! Посмотри в папку frontend/bindings/RangeLog
import { DatabaseService } from "../bindings/RangeLog/services"; 
import { SQLite } from "../bindings/RangeLog/services/models"; 

function App() {
  // 1. Создаем состояние для массива логов
  const [logs, setLogs] = useState<SQLite[]>([]);

  // 2. Функция загрузки данных
    const loadData = async () => {
    try {
        const data = await DatabaseService.GetRangeLog();
        // Если data равно null, ставим пустой массив []
        setLogs(data || []); 
    } catch (e) {
        console.error("Ошибка при загрузке данных:", e);
    }
  }

  // 3. Вызываем загрузку один раз при старте
  useEffect(() => {
    loadData();
  }, []);

  // 4. Отрисовка интерфейса
  return (
    <div style={{ padding: '20px', fontFamily: 'sans-serif' }}>
      <h1>RangeLog</h1>
      <button onClick={loadData}>Обновить данные</button>
      
      <table style={{ marginTop: '20px', width: '100%', borderCollapse: 'collapse' }}>
        <thead>
          <tr>
            <th>ID</th>
            <th>Дата</th>
            <th>Время</th>
            <th>Расстояние</th>
            <th>Действие</th>
          </tr>
        </thead>
        <tbody>
          {logs.map((log) => (
            <tr key={log.Id}>
              <td>{log.Id}</td>
              <td>{log.Date}</td>
              <td>{log.Time}</td>
              <td>{log.Distance} см</td>
              <td>{log.Action}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default App