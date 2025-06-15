CREATE DATABASE IF EXISTS db_cartshop;

USE db_cartshop;

-- ตารางเก็บข้อมูลผู้ใช้ (User)
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสผู้ใช้ (PK)',
    username VARCHAR(100) NULL UNIQUE COMMENT 'ชื่อผู้ใช้ (unique)',
    email VARCHAR(255) NULL UNIQUE COMMENT 'อีเมลผู้ใช้ (unique)',
    password VARCHAR(255) NULL COMMENT 'รหัสผ่าน (hashed)',
    is_active TINYINT NULL DEFAULT 0 COMMENT '0 = เปิดใช้งาน, 1 = ปิดใช้งาน',
    cre_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่สร้างข้อมูล',
    upd_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่แก้ไขล่าสุด',
    prog_id VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่ดำเนินการ/ตรวจสอบ',
    cre_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่สร้างข้อมูล',
    upd_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่แก้ไขข้อมูลล่าสุด'
) COMMENT = 'ตารางเก็บข้อมูลผู้ใช้ระบบ';

-- ตารางเก็บข้อมูลสินค้า (Product)
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสสินค้า (PK)',
    sku VARCHAR(50) NULL COMMENT 'sku สินค้า',
    name VARCHAR(255) NULL COMMENT 'ชื่อสินค้า',
    description TEXT COMMENT 'รายละเอียดสินค้า',
    price DECIMAL(10, 2) NULL COMMENT 'ราคาต่อหน่วย',
    stock INT DEFAULT 0 COMMENT 'จำนวนสินค้าคงเหลือในสต็อก',
    is_active TINYINT NULL DEFAULT 0 COMMENT '0 = เปิดใช้งาน, 1 = ปิดใช้งาน',
    cre_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่สร้างข้อมูล',
    upd_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่แก้ไขล่าสุด',
    cre_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่สร้างสินค้า',
    upd_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่แก้ไขสินค้าล่าสุด'
) COMMENT = 'ตารางเก็บข้อมูลสินค้า';

-- ตารางเก็บตะกร้าสินค้า (Cart)
-- ตะกร้าหนึ่งใบต่อผู้ใช้ เก็บสถานะชั่วคราวก่อนสั่งซื้อ
CREATE TABLE carts (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสตะกร้า (PK)',
    user_id INT NULL COMMENT 'รหัสผู้ใช้เจ้าของตะกร้า (FK ไป users.id)',
    is_active TINYINT NULL DEFAULT 0 COMMENT '0 = เปิดใช้งาน, 1 = ปิดใช้งาน',
    cre_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่สร้างข้อมูล',
    upd_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่แก้ไขล่าสุด',
    prog_id VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่ดำเนินการ/ตรวจสอบ',
    cre_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่สร้างตะกร้า',
    upd_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่แก้ไขตะกร้าล่าสุด',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) COMMENT = 'ตารางเก็บตะกร้าสินค้าของผู้ใช้';

-- ตารางเก็บรายการสินค้าในตะกร้า (Cart Items)
CREATE TABLE cart_items (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสรายการในตะกร้า (PK)',
    cart_id INT NULL COMMENT 'รหัสตะกร้า (FK ไป carts.id)',
    product_id INT NULL COMMENT 'รหัสสินค้า (FK ไป products.id)',
    quantity INT NULL DEFAULT 1 COMMENT 'จำนวนสินค้าที่ใส่ในตะกร้า',
    added_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่เพิ่มสินค้าในตะกร้า',
    is_active TINYINT NULL DEFAULT 0 COMMENT '0 = เปิดใช้งาน, 1 = ปิดใช้งาน',
    cre_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่สร้างข้อมูล',
    upd_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่แก้ไขล่าสุด',
    prog_id VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่ดำเนินการ/ตรวจสอบ',
    cre_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่สร้างคูปอง',
    upd_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่แก้ไขคูปองล่าสุด',
    FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id)
) COMMENT = 'รายการสินค้าภายในตะกร้าแต่ละใบ';

-- ตารางเก็บที่อยู่จัดส่ง (Shipping Address)
CREATE TABLE addresses (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสที่อยู่จัดส่ง (PK)',
    user_id INT NULL COMMENT 'รหัสผู้ใช้เจ้าของที่อยู่ (FK ไป users.id)',
    full_name VARCHAR(255) NULL COMMENT 'ชื่อเต็มผู้รับสินค้า',
    phone VARCHAR(20) COMMENT 'เบอร์โทรผู้รับสินค้า',
    address_line1 VARCHAR(255) NULL COMMENT 'ที่อยู่บรรทัดที่ 1',
    address_line2 VARCHAR(255) COMMENT 'ที่อยู่บรรทัดที่ 2 (ถ้ามี)',
    city VARCHAR(100) NULL COMMENT 'เมือง',
    state VARCHAR(100) COMMENT 'รัฐหรือจังหวัด',
    postal_code VARCHAR(20) COMMENT 'รหัสไปรษณีย์',
    country VARCHAR(100) NULL COMMENT 'ประเทศ',
    is_active TINYINT NULL DEFAULT 0 COMMENT '0 = เปิดใช้งาน, 1 = ปิดใช้งาน',
    cre_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่สร้างข้อมูล',
    upd_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่แก้ไขล่าสุด',
    prog_id VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่ดำเนินการ/ตรวจสอบ',
    cre_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่สร้างที่อยู่',
    upd_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่แก้ไขที่อยู่ล่าสุด',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) COMMENT = 'ตารางเก็บข้อมูลที่อยู่จัดส่งของผู้ใช้';

-- ตารางคำสั่งซื้อ (Orders)
CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสคำสั่งซื้อ (PK)',
    user_id INT NULL COMMENT 'รหัสผู้ใช้เจ้าของคำสั่งซื้อ (FK ไป users.id)',
    address_id INT NULL COMMENT 'รหัสที่อยู่จัดส่งในคำสั่งซื้อ (FK ไป addresses.id)',
    total_price DECIMAL(10, 2) NULL COMMENT 'ราคารวมทั้งหมดของคำสั่งซื้อ',
    status ENUM(
        'pending',
        'paid',
        'shipped',
        'completed',
        'cancelled'
    ) DEFAULT 'pending' COMMENT 'สถานะคำสั่งซื้อ',
    is_active TINYINT NULL DEFAULT 0 COMMENT '0 = เปิดใช้งาน, 1 = ปิดใช้งาน',
    cre_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่สร้างข้อมูล',
    upd_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่แก้ไขล่าสุด',
    prog_id VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่ดำเนินการ/ตรวจสอบ',
    cre_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่สร้างคำสั่งซื้อ',
    upd_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่แก้ไขคำสั่งซื้อ',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (address_id) REFERENCES addresses(id)
) COMMENT = 'ตารางเก็บคำสั่งซื้อของผู้ใช้';

-- ตารางรายการสินค้าภายในคำสั่งซื้อ (Order Items)
CREATE TABLE order_items (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสรายการสินค้าคำสั่งซื้อ (PK)',
    order_id INT NULL COMMENT 'รหัสคำสั่งซื้อ (FK ไป orders.id)',
    product_id INT NULL COMMENT 'รหัสสินค้า (FK ไป products.id)',
    quantity INT NULL COMMENT 'จำนวนสินค้าที่สั่ง',
    price DECIMAL(10, 2) NULL COMMENT 'ราคาต่อหน่วย ณ เวลาสั่งซื้อ (สำคัญเก็บเพื่อประวัติ)',
    is_active TINYINT NULL DEFAULT 0 COMMENT '0 = เปิดใช้งาน, 1 = ปิดใช้งาน',
    cre_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่สร้างข้อมูล',
    upd_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่แก้ไขล่าสุด',
    prog_id VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่ดำเนินการ/ตรวจสอบ',
    cre_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่สร้างคูปอง',
    upd_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่แก้ไขคูปองล่าสุด',
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id)
) COMMENT = 'รายการสินค้าของแต่ละคำสั่งซื้อ';

-- ตารางคูปองส่วนลด (Coupon / Promotion) - Optional
CREATE TABLE coupons (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสคูปอง (PK)',
    code VARCHAR(50) NULL UNIQUE COMMENT 'โค้ดคูปอง',
    description TEXT COMMENT 'รายละเอียดคูปอง',
    discount_percent INT DEFAULT 0 COMMENT 'ส่วนลดแบบเปอร์เซ็นต์',
    discount_amount DECIMAL(10, 2) DEFAULT 0 COMMENT 'ส่วนลดแบบจำนวนเงิน',
    active BOOLEAN DEFAULT TRUE COMMENT 'สถานะใช้งานคูปอง',
    expiry_date DATE COMMENT 'วันหมดอายุคูปอง',
    is_active TINYINT NULL DEFAULT 0 COMMENT '0 = เปิดใช้งาน, 1 = ปิดใช้งาน',
    cre_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่สร้างข้อมูล',
    upd_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่แก้ไขล่าสุด',
    prog_id VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่ดำเนินการ/ตรวจสอบ',
    cre_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่สร้างคูปอง',
    upd_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่แก้ไขคูปองล่าสุด'
) COMMENT = 'ตารางเก็บคูปองส่วนลด';

-- ตารางเก็บการใช้คูปองกับคำสั่งซื้อ (Order Coupons) - Optional
CREATE TABLE order_coupons (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสการใช้คูปอง (PK)',
    order_id INT NULL COMMENT 'รหัสคำสั่งซื้อ (FK ไป orders.id)',
    coupon_id INT NULL COMMENT 'รหัสคูปองที่ใช้ (FK ไป coupons.id)',
    discount_amount DECIMAL(10, 2) NULL COMMENT 'จำนวนส่วนลดที่ใช้ในคำสั่งซื้อ',
    is_active TINYINT NULL DEFAULT 0 COMMENT '0 = เปิดใช้งาน, 1 = ปิดใช้งาน',
    cre_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่สร้างข้อมูล',
    upd_by VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่แก้ไขล่าสุด',
    prog_id VARCHAR(50) DEFAULT NULL COMMENT 'รหัสผู้ใช้ที่ดำเนินการ/ตรวจสอบ',
    cre_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่สร้างคูปอง',
    upd_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่แก้ไขคูปองล่าสุด',
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (coupon_id) REFERENCES coupons(id)
) COMMENT = 'ตารางเก็บข้อมูลการใช้คูปองในคำสั่งซื้อ';