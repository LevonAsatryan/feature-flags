import { Injectable } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { User } from './entities/user.entity';
import { UserRole } from 'src/shared/enums/user-role.enum';
import { randomPasswordGenerator } from 'src/shared/util';
import { Company } from 'src/company/entities/company.entity';

@Injectable()
export class UsersService {
  constructor(
    @InjectRepository(User) private userRepository: Repository<User>
  ) {}

  createDefaultAdminUser(email: string, company: Company) {
    const adminUser = new User();
    adminUser.email = email;
    adminUser.role = UserRole.ADMIN;
    adminUser.password = randomPasswordGenerator();
    adminUser.company = company;

    /**
     * TODO: Send an email to the user with the password
     */

    return this.userRepository.save(adminUser);
  }

  create(createUserDto: CreateUserDto) {
    return 'This action adds a new user';
  }

  findAll() {
    return `This action returns all users`;
  }

  findOne(id: number) {
    return `This action returns a #${id} user`;
  }

  update(id: number, updateUserDto: UpdateUserDto) {
    return `This action updates a #${id} user`;
  }

  remove(id: number) {
    return `This action removes a #${id} user`;
  }
}
