<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\KategoriArtikelController;
use App\Http\Controllers\ArtikelController;
use App\Http\Controllers\CutiDokterController;
use App\Http\Controllers\FAQController;
use App\Http\Controllers\AuthController;
use App\Http\Controllers\UlasanController;
use App\Http\Controllers\ObatController;
use App\Http\Controllers\AntrianController;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "web" middleware group. Make something great!
|
*/

Route::get('/', function () {
    return view('dashboard');
});

Route::get('/login/form', [AuthController::class, 'showLoginForm'])->name('loginForm');
Route::post('/login', [AuthController::class, 'login']) ->name("login");
Route::get('/register/form', [AuthController::class, 'showRegisterForm'])->name('registerForm');
Route::post('/register', [AuthController::class, 'register'])->name("register");
Route::post('/logout', [AuthController::class, 'logout'])->name('logout');

Route::get('/kategori-artikel', [KategoriArtikelController::class, 'index'])->name('kategori-artikel.index');
Route::get('/kategori-artikel/create', [KategoriArtikelController::class, 'create'])->name('kategori-artikel.create');
Route::post('/kategori-artikel', [KategoriArtikelController::class, 'store'])->name('kategori-artikel.store');
Route::get('/kategori-artikel/{id}/edit', [KategoriArtikelController::class, 'edit'])->name('kategori-artikel.edit');
Route::put('/kategori-artikel/{id}', [KategoriArtikelController::class, 'update'])->name('kategori-artikel.update');
Route::delete('/kategori-artikel/{id}', [KategoriArtikelController::class, 'destroy'])->name('kategori-artikel.destroy');

Route::get('/artikels', [ArtikelController::class, 'index'])->name('artikel.index');
Route::get('/artikel/create', [ArtikelController::class, 'create'])->name('artikel.create');
Route::post('/artikel/store', [ArtikelController::class, 'store'])->name('artikel.store');
Route::get('/artikels/{id}', [ArtikelController::class, 'show'])->name('artikel.show');
Route::get('/artikels/{id}/edit', [ArtikelController::class, 'edit'])->name('artikel.edit');
Route::put('/artikels/{id}', [ArtikelController::class, 'update'])->name('artikel.update');
Route::delete('/artikels/{id}', [ArtikelController::class, 'destroy'])->name('artikel.destroy');

Route::get('/cutidokter', [CutiDokterController::class, 'index'])->name('cuti-dokter.index');
Route::get('/cuti-dokter/create', [CutiDokterController::class, 'create'])->name('cuti-dokter.create');
Route::post('/cuti-dokter', [CutiDokterController::class, 'store'])->name('cuti-dokter.store');
Route::get('/cuti-dokter/{id}/edit', [CutiDokterController::class, 'edit'])->name('cuti-dokter.edit');
Route::put('/cuti-dokter/{id}', [CutiDokterController::class, 'update'])->name('cuti-dokter.update');
Route::delete('/cuti-dokter/{id}', [CutiDokterController::class, 'destroy'])->name('cuti-dokter.destroy');

Route::get('/all/faqs', [FAQController::class, 'index'])->name('faq.index');
Route::get('/faq/create', [FAQController::class, 'create'])->name('faq.create');
Route::post('/faq', [FAQController::class, 'store'])->name('faq.store');
Route::get('/faq/{id}/edit', [FAQController::class, 'edit'])->name('faq.edit');
Route::put('/faq/{id}', [FAQController::class, 'update'])->name('faq.update');
Route::delete('/faq/{id}', [FAQController::class, 'destroy'])->name('faq.destroy');

Route::get('/ulasan', [UlasanController::class, 'index'])->name('ulasan.index');
Route::get('/ulasan/create', [UlasanController::class, 'create'])->name('ulasan.create');
Route::post('/ulasan/store', [UlasanController::class, 'store'])->name('ulasan.store');
Route::get('/ulasan/{id}/edit', [UlasanController::class, 'edit'])->name('ulasan.edit');
Route::put('/ulasan/{id}', [UlasanController::class, 'update'])->name('ulasan.update');
Route::delete('/ulasan/{id}', [UlasanController::class, 'destroy'])->name('ulasan.destroy');

Route::get('/obat', [ObatController::class, 'index'])->name('obat.index');
Route::get('/obat/create', [ObatController::class, 'create'])->name('obat.create');
Route::post('/obat/store', [ObatController::class, 'store'])->name('obat.store');
Route::get('/obat/{id}/edit', [ObatController::class, 'edit'])->name('obat.edit');
Route::put('/obat/{id}/update', [ObatController::class, 'update'])->name('obat.update');
Route::delete('/obat/{id}/delete', [ObatController::class, 'destroy'])->name('obat.delete');


Route::get('/antrian/create', [AntrianController::class, 'create'])->name('antrian.create');
Route::post('/antrian/store', [AntrianController::class, 'store'])->name('antrian.store');
Route::get('/antrian', [AntrianController::class, 'index'])->name('antrian.index');
Route::delete('/antrian/{id}', [AntrianController::class, 'destroy'])->name('antrian.destroy');
