from aiogram import Bot, Dispatcher, types
from aiogram.contrib.fsm_storage.memory import MemoryStorage
from aiogram.utils import executor
from aiogram.dispatcher import FSMContext
from aiogram.dispatcher.filters.state import State, StatesGroup
import requests

bot = Bot(token='5819456617:AAHJLx8uAzxmlgmVjN5ekDU9V-AcSdoFRSw')
storage = MemoryStorage()
dp = Dispatcher(bot, storage=storage)

class LoginForm(StatesGroup):
    login = State()
    password = State()
    main = State()

@dp.message_handler(commands=['start', 'login'])
async def start_command(message: types.Message):
    await message.reply("Добро пожаловать! Чтобы Войти, введите ваше имя.")
    await LoginForm.login.set()

@dp.message_handler(state=LoginForm.login)
async def process_login(message: types.Message, state: FSMContext):
    login = message.text
    await state.update_data(login = login)
    await message.reply("Введите пароль")
    await LoginForm.password.set()

@dp.message_handler(state=LoginForm.password)
async def process_password(message: types.Message, state: FSMContext):
    password = message.text
    await state.update_data(password = password)

    data = await state.get_data()
    username = data.get('login')
    password = data.get('password')

    payload = {
        "username": username,
        "password": password
    }

    response = requests.post("http://127.0.0.1:8000/Login", json=payload)

    if response.status_code == 200:
        await message.reply("Вы успешно вошли!")
        await state.update_data(token = response.text)
    else:
        await message.reply("Что-то пошло не так брат!")

    await LoginForm.main.set()