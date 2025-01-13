async function getLevel1Web2Answer() {
  const url = "https://raw.githubusercontent.com/nepp-tumsat/nepp-ctf/refs/heads/main/questions/level1/web/2/level1-web-2.json";
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`レスポンスステータス: ${response.status}`);
    }
    await response.json();
  } catch (error) {
    console.error(error.message);
  }
}

getLevel1Web2Answer();